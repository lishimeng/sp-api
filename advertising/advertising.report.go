package advertising

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/rest"
	"io"
)

type ReportCategory string
type ReportType string

const (
	Brands  ReportCategory = "brands"
	Display ReportCategory = "display"
)

const (
	Campaign  ReportType = "Campaign"
	Placement ReportType = "Placement"
	AdGroup   ReportType = "Ad group"
)

type DateDuration struct {
	StartDate string `json:"startDate"` // yyyy-mm-dd
	EndDate   string `json:"endDate"`   // yyyy-mm-dd
}

type RptFilter struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

type RptConfig struct {
	AdProduct    string      `json:"adProduct"`
	GroupBy      []string    `json:"groupBy"`
	Columns      []string    `json:"columns"`
	ReportTypeId ReportType  `json:"reportTypeId"`
	TimeUnit     string      `json:"timeUnit"`
	Format       string      `json:"format"`
	Filters      []RptFilter `json:"filters"`
}

type RptHeader struct {
	Name string `json:"name"`
	DateDuration
}

type RptStatus struct {
	ReportId      string `json:"reportId"`
	Status        string `json:"status"`
	FailureReason string `json:"failureReason"`

	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	GeneratedAt string `json:"generatedAt"`

	FileSize     string `json:"fileSize"`
	Url          string `json:"url"`
	UrlExpiresAt string `json:"urlExpiresAt"`
}

type RptReq struct {
	RptHeader
	Config RptConfig `json:"configuration"`
}

func (c *Client) CreateReport(req RptReq) (err error) {
	var action = "/reporting/reports"
	err = c.request().Path(action).Json(rest.POST)
	return
}

func (c *Client) ReportStatus(id string) (err error) {
	var action = "/reporting/reports"
	var resp RptStatus
	err = c.request().Path(action).Path(id).Response(&resp).Json(rest.GET)
	if err != nil {
		log.Info(err)
		return
	}
	bs, _ := json.Marshal(resp)
	log.Info(string(bs))
	return
}

func (c *Client) DownloadReport(url string, h func(size int64, reader io.Reader)) (err error) {
	err = rest.NewRequest(url, c.ssl).Download(h)
	return
}
