package data

const (
	EURUSD = "EURUSD"
)

type FullForexQuoteBulkData struct {
	Docs *[]FullForexQuote `json:"docs"`
}

type HistoricalDailyForexBulkData struct {
	Docs *[]HistoricalDailyForex `json:"docs"`
}

type IntraDayForexBulkData struct {
	Docs *[]IntraDayForex `json:"docs"`
}
