package seller

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/rest"
	"net/http"
	"time"
)

type ReportFilter struct {
	params *CreateReportSpecification
}

type ReportFilterFunc func(filter *CreateReportSpecification)

var WithReportType = func(rt ReportType) ReportFilterFunc {
	return func(filter *CreateReportSpecification) {
		filter.ReportType = rt
	}
}

var WithMarketPlace = func(mp common.Marketplace) ReportFilterFunc {
	return func(filter *CreateReportSpecification) {
		_, c := common.GetCentralURL(mp)
		filter.MarketplaceIds = append(filter.MarketplaceIds, c.Id)
	}
}

// WithDuration 设置要查询的时间段
var WithDuration = func(start, end time.Time) ReportFilterFunc {
	return func(filter *CreateReportSpecification) {
		filter.DataStartTime = start.Format(time.RFC3339)
		filter.DataEndTime = end.Format(time.RFC3339)
	}
}

// CreateReport 创建报告
func (c *Client) CreateReport(filterFunc ...ReportFilterFunc) (reportId string, err error) {

	var action = "/reports/2021-06-30/reports"

	var filter = &CreateReportSpecification{}
	var result CreateReportResponse
	for _, f := range filterFunc {
		if f != nil {
			f(filter)
		}
	}

	bs, err := json.Marshal(filter)
	if err != nil {
		return
	}
	log.Info("CreateReport: %s", string(bs))

	err = c.sellerRequest().Path(action).
		Accept("application/json").
		Expect(http.StatusAccepted).
		Body(filter).
		Response(&result).Json(rest.POST)

	if err != nil {
		return
	}

	bs, err = json.Marshal(result)
	if err != nil {
		return
	}

	log.Info("result: %s", string(bs))

	reportId = result.ReportId

	return
}

type ReportResponse struct {
	Errors []ErrorResponse `json:"errors,omitempty"`
	ReportPayload
}

type ReportProcessInfo struct {
	ProcessingStatus    string `json:"processingStatus,omitempty"`
	ProcessingStartTime string `json:"processingStartTime,omitempty"`
	ProcessingEndTime   string `json:"processingEndTime,omitempty"`
}

type ReportBaseInfo struct {
	ReportId         string   `json:"reportId,omitempty"`
	MarketplaceIds   []string `json:"marketplaceIds,omitempty"`
	ReportDocumentId string   `json:"reportDocumentId,omitempty"`
	ReportType       string   `json:"reportType,omitempty"`
}

type ReportPayload struct {
	ReportBaseInfo
	DataStartTime string `json:"dataStartTime,omitempty"`
	DataEndTime   string `json:"dataEndTime,omitempty"`
	CreatedTime   string `json:"createdTime,omitempty"`
	ReportProcessInfo
}

func (c *Client) GetReport(reportId string) (payload ReportPayload, err error) {
	var action = "/reports/2021-06-30/reports"
	var result ReportResponse
	err = c.sellerRequest().Path(action).Path(reportId).
		Accept("application/json").
		Expect(http.StatusOK).Response(&result).Json(rest.GET)

	if err != nil {
		return
	}
	payload = result.ReportPayload
	return
}

type ReportDocumentResp struct {
	ReportDocumentId     string `json:"reportDocumentId,omitempty"`
	Url                  string `json:"url,omitempty"`
	CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`
}

func (c *Client) GetReportDocument(reportDocumentId string) (payload ReportDocumentResp, err error) {
	var action = "/reports/2021-06-30/reports/documents"
	var result ReportDocumentResp
	err = c.sellerRequest().Path(action).Path(reportDocumentId).
		Accept("application/json").
		Expect(http.StatusOK).Response(&result).Json(rest.GET)

	if err != nil {
		return
	}
	payload = result
	return
}
