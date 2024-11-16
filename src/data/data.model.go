package data

type FullForexQuote struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Price                float32 `json:"price"`
	ChangesPercentage    float32 `json:"changesPercentage"`
	Change               float32 `json:"change"`
	DayLow               float32 `json:"dayLow"`
	DayHigh              float32 `json:"dayHigh"`
	YearLow              float32 `json:"yearLow"`
	YearHigh             float32 `json:"yearHigh"`
	MarketCap            float32 `json:"marketCap"`
	PriceAvg50           float32 `json:"priceAvg50"`
	PriceAvg200          float32 `json:"priceAvg200"`
	Volume               int32   `json:"volume"`
	AvgVolume            float32 `json:"avgVolume"`
	Open                 float32 `json:"open"`
	PreviousClose        float32 `json:"previousClose"`
	EarningsAnnouncement string  `json:"earningsAnnouncement"`
	SharesOutstanding    float32 `json:"sharesOutstanding"`
	Timestamp            int64   `json:"timestamp"`
}

type IntraDayForex struct {
	Date      string  `json:"date"`
	Open      float32 `json:"open"`
	Low       float32 `json:"low"`
	High      float32 `json:"high"`
	Close     float32 `json:"close"`
	Volume    float32 `json:"volume"`
	Timestamp int64   `json:"timestamp"`
}

type HistoricalForex struct {
	ID               string  `json:"_id"`
	Symbol           string  `json:"symbol"`
	Date             string  `json:"date"`
	Open             float32 `json:"open"`
	High             float32 `json:"high"`
	Low              float32 `json:"low"`
	Close            float32 `json:"close"`
	AdjClose         float32 `json:"adjClose"`
	Volume           float32 `json:"volume"`
	UnadjustedVolume float32 `json:"unadjustedVolume"`
	Change           float32 `json:"change"`
	ChangePercent    float32 `json:"changePercent"`
	Vwap             float32 `json:"vwap"`
	Label            string  `json:"label"`
	ChangeOverTime   float32 `json:"changeOverTime"`
	Timestamp        int64   `json:"timestamp"`
}

type HistoricalDailyForex struct {
	Symbol     string            `json:"symbol"`
	Historical []HistoricalForex `json:"historical"`
}
