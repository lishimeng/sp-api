package seller

type ReportType string

const (

	// Brand Analytics

	AnalyticsBrand          ReportType = "GET_BRAND_ANALYTICS_MARKET_BASKET_REPORT"   //购物车 分析报告
	AnalyticsSearchKey      ReportType = "GET_BRAND_ANALYTICS_SEARCH_TERMS_REPORT"    // 搜索词分析报告
	AnalyticsRepeatPurchase ReportType = "GET_BRAND_ANALYTICS_REPEAT_PURCHASE_REPORT" // 复购报告
)
