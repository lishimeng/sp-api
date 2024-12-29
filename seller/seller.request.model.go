package seller

type MoneyType struct {
	CurrencyCode string  `json:"CurrencyCode,omitempty"`
	Amount       float64 `json:"Amount,omitempty"`
}

type Points struct {
	PointsNumber        int       `json:"PointsNumber,omitempty"`
	PointsMonetaryValue MoneyType `json:"PointsMonetaryValue,omitempty"`
}

// CreateReportSpecification 业务请求:创建报告
type CreateReportSpecification struct {
	ReportOptions  ReportOptions `json:"reportOptions,omitempty"`
	ReportType     ReportType    `json:"reportType"`
	DataStartTime  string        `json:"dataStartTime,omitempty"` // ISO 8601
	DataEndTime    string        `json:"dataEndTime,omitempty"`   // ISO 8601
	MarketplaceIds []string      `json:"marketplaceIds"`
}

type ReportOptions struct {
}

// CreateReportResponse 业务响应:创建报告
type CreateReportResponse struct {
	Errors   []ErrorResponse `json:"errors,omitempty"`
	ReportId string          `json:"reportId,omitempty"`
}

type ItemSearchResults struct {
	ErrorResponseWrapper
	NumberOfResults int64       `json:"numberOfResults,omitempty"`
	Pagination      Pagination  `json:"pagination"`
	Refinements     Refinements `json:"refinements"`
	Items           []Item      `json:"items,omitempty"`
}

type Pagination struct {
	NextToken     string `json:"nextToken,omitempty"`
	PreviousToken string `json:"previousToken,omitempty"`
}

type Refinements struct {
	Brands          []BrandRefinement          `json:"brands,omitempty"`
	Classifications []ClassificationRefinement `json:"classifications,omitempty"`
}

type Item struct {
	Asin string `json:"asin,omitempty"`
}

type BrandRefinement struct {
	NumberOfResults int64  `json:"numberOfResults,omitempty"`
	BrandName       string `json:"brandName,omitempty"`
}

type ClassificationRefinement struct {
	NumberOfResults  int64  `json:"numberOfResults,omitempty"`
	DisplayName      string `json:"displayName,omitempty"`
	ClassificationId string `json:"classificationId,omitempty"`
}

// ------------------------------------------------

// ErrorResponse 业务响应:错误响应
type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

type ErrorResponseWrapper struct {
	Errors []ErrorResponse `json:"errors,omitempty"`
}
