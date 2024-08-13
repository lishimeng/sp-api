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

	c.refreshAccessToken()
	err = rest.NewRequest(c.endPoint, c.ssl).
		Path(action).
		Authorization(c.tokenTemp.AccessToken).
		Header(rest.HeaderUserAgent, c.userAgent).
		Header("Host", c.endPoint).
		Accept("application/json").
		RequestTime(time.Now()).
		Expect(http.StatusAccepted).
		Body(filter).
		Response(&result).Json()

	if err != nil {
		return
	}

	reportId = result.ReportId

	return
}
