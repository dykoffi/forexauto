package db

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/dykoffi/forexauto/src/data"
)

func BenchmarkAddData(b *testing.B) {
	b.ReportAllocs()
	data, err := data.New().GetHistoricalDailyForex("EURUSD")

	docs := map[string]interface{}{
		"docs": data,
	}

	if err != nil {
		return
	}

	dataByte, err := json.Marshal(&docs)

	reader := bytes.NewBuffer(dataByte)

	if err != nil {
		return
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New().AddData(reader)
		time.Sleep(1 * time.Millisecond)
	}
}
