package data

import (
	"bytes"
	"encoding/json"
	"io"
)

func TransformToReader[T []FullForexQuote | []IntraDayForex | []HistoricalForex | FullForexQuoteBulkData | HistoricalDailyForexBulkData | IntraDayForexBulkData](data *T) (io.Reader, error) {

	dataByte, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(dataByte), nil
}
