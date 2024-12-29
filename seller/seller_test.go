package seller

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/lwa"
	"os"
	"testing"
	"time"
)

type TokenFactory struct {
	Seller lwa.AccessToken `json:"seller,omitempty"`
	Ads    lwa.AccessToken `json:"ads,omitempty"`
}

var tokenFile string

var (
	rt     = ""
	appid  = ""
	secret = ""
)
var ft TokenFactory

const (
	SellerType = "SELLER_"
	AdsType    = "ADS_"
)

var apiType = ""

func loadEnv() {
	tokenFile = "tmp_token.json"
	rt = os.Getenv(apiType + "AWS_REFRESH_KEY")

	appid = os.Getenv(apiType + "AWS_APP_ID")
	secret = os.Getenv(apiType + "AWS_SECRET")
}

func saveToken(t *testing.T, at lwa.AccessToken) {
	f, err := os.OpenFile(tokenFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Log(err)
		return
	}
	defer func() {
		_ = f.Close()
	}()
	if apiType == "SELLER_" {
		ft.Seller = at
	} else if apiType == "ADS_" {
		ft.Ads = at
	}
	bs, err := json.Marshal(ft)
	_, err = f.WriteString(string(bs))
	if err != nil {
		t.Log("保持token失败:" + err.Error())
		t.Log(err)
		return
	}
}

func loadToken(t *testing.T) (tf TokenFactory) {
	bs, err := os.ReadFile(tokenFile)
	if err != nil {
		t.Log("token文件不存在:" + tokenFile)
		return
	}
	err = json.Unmarshal(bs, &tf)
	if err != nil {
		t.Log("token内容错误:")
		t.Log(err)
		return
	}
	t.Log("reuse toke.")
	return
}

func createLwa() (*lwa.Client, string) {
	ua := UserAgent(appid, "v1.0")
	log.Info("init sdk: %s[%s]", appid, secret)
	connector := lwa.New(appid, secret, rt)
	return connector, ua
}

func createApi(t *testing.T) (c *Client) {
	loadEnv() // 初始化token

	at := loadToken(t)

	//seller.AutoRefreshToken = true
	connector, ua := createLwa()

	c = New(WithLwa(connector, func(token lwa.AccessToken) {
		saveToken(t, token) // 保存token
	}), WithTokenErrorListener(func(clientId string, err error) {
		t.Log("token error:", err)
	}), WithMarketplace(common.Us, true),
		WithLwaToken(at.Ads),
		WithUserAgent(ua))
	return
}

func TestSellerCategory(t *testing.T) {

	apiType = SellerType
	c := createApi(t)

	items, err := c.GetItems()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(items.NumberOfResults)
	t.Log(items.Pagination.PreviousToken, items.Pagination.NextToken)
}

func TestSellerRpt(t *testing.T) {

	apiType = SellerType
	c := createApi(t)

	rptFrom, _ := time.Parse(time.DateOnly, "2024-12-01")
	rptEnd, _ := time.Parse(time.DateOnly, "2024-12-15")

	reportId, err := c.CreateReport(WithReportType(AnalyticsBrand),
		WithMarketPlace(common.Us),
		WithDuration(rptFrom, rptEnd))

	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(reportId)
}

func TestGetReport(t *testing.T) {

	apiType = SellerType
	var rptId = "50427020085"
	c := createApi(t)
	payload, err := c.GetReport(rptId)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(payload)
	t.Logf("report status:%s", payload.ProcessingStatus)
}

func TestGetReportDocument(t *testing.T) {
	apiType = SellerType
	var docId = "amzn1.spdoc.1.4.na.33dfc2ef-16d6-48ad-a7db-9eedc79632e9.T211K39FNDS4D7.88700"
	c := createApi(t)
	payload, err := c.GetReport(docId)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(payload)
}
