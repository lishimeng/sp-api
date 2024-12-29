package seller

import (
	"encoding/json"
	"github.com/lishimeng/sp-api/lwa"
	"testing"
)

func TestCreateAdsCode(t *testing.T) {

	apiType = AdsType

	loadEnv() // 初始化token

	conn, ua := createLwa()
	t.Logf("user-agent:%s", ua)
	uri, err := conn.Code("https://amazon.com", lwa.ScopeAdvertisingCampaign)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("uri:%s", uri)
}

func TestTokenWithCode(t *testing.T) {
	apiType = AdsType
	var code = ""
	var redirectUri = "https://amazon.com"
	conn, ua := createLwa()
	t.Logf("user-agent:%s", ua)
	at, err := conn.GetAccessTokenWithCode(code, redirectUri)
	if err != nil {
		t.Fatal(err)
		return
	}
	bs, _ := json.Marshal(at)
	t.Logf("token:")
	t.Log(string(bs))
}

func TestGetAdsReport(t *testing.T) {

	apiType = AdsType
	var err error
	req := AdsRptReq{
		Name: "sp_campaigns_report[12/5-7/10]",
		DateRange: DateRange{
			StartDate: "2024-12-05",
			EndDate:   "2024-12-10",
		},
		Configuration: AdsConfig{
			ReportTypeId: AdsRptCampaign,
			TimeUnit:     Daily,
			Format:       "GZIP_JSON",
		},
	}

	c := createApi(t)
	_, err = c.GetAdsReport(req)
	if err != nil {
		t.Fatal(err)
		return
	}
}
