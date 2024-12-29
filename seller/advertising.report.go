package seller

import (
	"github.com/lishimeng/sp-api/rest"
	"net/http"
)

type AdsRptType string

const (
	AdsRptCampaign AdsRptType = "spCampaigns"
)

type AdsRptReq struct {
	Name string `json:"name,omitempty"`
	DateRange
	Configuration AdsConfig `json:"configuration"`
}

type AdsConfig struct {
	AdProduct    string      `json:"adProduct,omitempty"`
	ReportTypeId AdsRptType  `json:"reportTypeId,omitempty"`
	GroupBy      []string    `json:"groupBy,omitempty"`
	Columns      []string    `json:"columns,omitempty"`
	TimeUnit     TimeUnit    `json:"timeUnit,omitempty"`
	Format       string      `json:"format,omitempty"`
	Filters      []AdsFilter `json:"filters,omitempty"`
}
type AdsFilter struct {
	Field  string   `json:"field,omitempty"`
	Values []string `json:"values,omitempty"`
}

type TimeUnit string

const (
	Summary TimeUnit = "SUMMARY"
	Daily   TimeUnit = "DAILY"
)

type DateRange struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

func (c *Client) GetAdsReport(req AdsRptReq) (payload ReportPayload, err error) {
	var action = "/reporting/reports"
	var result map[string]any
	err = c.adsRequest().Path(action).
		Accept("application/json").
		Body(req).
		Expect(http.StatusOK).Response(&result).Json(rest.POST)

	if err != nil {
		return
	}

	return
}
