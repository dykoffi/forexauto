package data

type FullForexQuote struct {
	Symbol               string
	Name                 string
	Price                float32
	ChangesPercentage    float32
	Change               float32
	DayLow               float32
	DayHigh              float32
	YearLow              float32
	YearHigh             float32
	MarketCap            float32
	PriceAvg50           float32
	PriceAvg200          float32
	Volume               int32
	AvgVolume            float32
	Open                 float32
	PreviousClose        float32
	EarningsAnnouncement string
	SharesOutstanding    float32
	Timestamp            int
}

type IntraDayForex struct {
	Date   string
	Open   float32
	Low    float32
	High   float32
	Close  float32
	Volume float32
}

type HistoricalForex struct {
	Date             string
	Open             float32
	High             float32
	Low              float32
	Close            float32
	AdjClose         float32
	Volume           float32
	UnadjustedVolume float32
	Change           float32
	ChangePercent    float32
	Vwap             float32
	Label            string
	ChangeOverTime   float32
}

type DailyForex struct {
	Symbol     string
	Historical []HistoricalForex
}
