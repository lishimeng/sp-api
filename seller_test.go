package sp_api

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/lwa"
	"github.com/lishimeng/sp-api/seller"
	"os"
	"testing"
	"time"
)

func saveToken(t *testing.T, at lwa.AccessToken) {
	f, err := os.OpenFile(tokenFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Log(err)
		return
	}
	defer func() {
		_ = f.Close()
	}()
	bs, err := json.Marshal(at)
	_, err = f.WriteString(string(bs))
	if err != nil {
		t.Log("保持token失败:" + err.Error())
		t.Log(err)
		return
	}
}

func loadToken(t *testing.T) (at lwa.AccessToken) {
	bs, err := os.ReadFile(tokenFile)
	if err != nil {
		t.Log("token文件不存在:" + tokenFile)
		return
	}
	err = json.Unmarshal(bs, &at)
	if err != nil {
		t.Log("token内容错误:")
		t.Log(err)
		return
	}
	t.Log("reuse toke.")
	return
}

func createLwa() (*lwa.Client, string) {
	ua := seller.UserAgent(appid, "v1.0")
	log.Info("init sdk: %s[%s]", appid, secret)
	connector := lwa.New(appid, secret, rt)
	return connector, ua
}

func createApi(t *testing.T) (c *seller.Client) {
	at := loadToken(t)

	//seller.AutoRefreshToken = true
	connector, ua := createLwa()

	c = seller.New(seller.WithLwa(connector, func(token lwa.AccessToken) {
		saveToken(t, token) // 保存token
	}), seller.WithMarketplace(common.Us, true),
		seller.WithLwaToken(at),
		seller.WithUserAgent(ua))
	return
}

func TestSellerCategory(t *testing.T) {

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

	c := createApi(t)

	rptFrom, _ := time.Parse(time.DateOnly, "2024-07-01")
	rptEnd, _ := time.Parse(time.DateOnly, "2024-07-30")

	reportId, err := c.CreateReport(seller.WithReportType(seller.AnalyticsBrand),
		seller.WithMarketPlace(common.Us),
		seller.WithDuration(rptFrom, rptEnd))

	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(reportId)
}
