package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	lop "github.com/samber/lo/parallel"
)

type DataInterface interface {
	New() *DataService
	GetFullForexQuote() (*[]FullForexQuote, error)
	GetIntraDayForex() (*[]IntraDayForex, error)
	GetHistoricalDailyForex() (*[]HistoricalForex, error)
}

type DataService struct {
	host      string
	apiKey    string
	symbol    string
	timeframe string
	config    *config.ConfigService
}

var (
	iDataService DataService
	once         sync.Once
)

func New(config *config.ConfigService) *DataService {

	once.Do(func() {
		iDataService = DataService{
			apiKey:    config.GetOrThrow("FOREX_API_KEY"),
			host:      config.GetOrThrow("FOREX_BASE_URL"),
			symbol:    config.GetOrThrow("FOREX_SYMBOL"),
			timeframe: config.GetOrThrow("FOREX_TIMEFRAME"),
		}
	})

	return &iDataService
}

func (ds *DataService) GetFullForexQuote() (*[]FullForexQuote, error) {

	path := fmt.Sprintf("/quote/%s", ds.symbol)
	finalUrl := fmt.Sprintf("%s/%s?apikey=%s", ds.host, path, ds.apiKey)
	res, err := http.Get(finalUrl)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return nil, err
	}

	var dataBody []FullForexQuote

	if err := json.Unmarshal(body, &dataBody); err != nil {
		return nil, err
	}

	fullForexQuoteData := lop.Map(dataBody, func(item FullForexQuote, index int) FullForexQuote {
		item.ID = fmt.Sprintf("%d", item.Timestamp)
		return item
	})

	return &fullForexQuoteData, nil
}

func (ds *DataService) GetIntraDayForex(from string, to string) (*[]IntraDayForex, error) {
	path := fmt.Sprintf("/historical-chart/%s/%s", ds.timeframe, ds.symbol)
	finalUrl := fmt.Sprintf("%s/%s?from=%s&to=%s&apikey=%s", ds.host, path, from, to, ds.apiKey)
	res, err := http.Get(finalUrl)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var dataBody []IntraDayForex

	if err := json.Unmarshal(body, &dataBody); err != nil {
		return nil, err
	}

	intraDayForexData := lop.Map(dataBody, func(item IntraDayForex, index int) IntraDayForex {
		tps, _ := time.Parse("2006-01-02 15:04:05", item.Date)
		item.ID = fmt.Sprintf("%d", tps.Unix())
		item.Timestamp = tps.Unix()
		return item
	})

	return &intraDayForexData, nil
}

func (ds *DataService) GetHistoricalDailyForex() (*[]HistoricalForex, error) {

	path := fmt.Sprintf("/historical-price-full/%s", ds.symbol)
	finalUrl := fmt.Sprintf("%s/%s?apikey=%s", ds.host, path, ds.apiKey)
	res, err := http.Get(finalUrl)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return nil, err
	}

	var dataBody HistoricalDailyForex

	if err := json.Unmarshal(body, &dataBody); err != nil {
		return nil, err
	}

	historicalData := lop.Map(dataBody.Historical, func(item HistoricalForex, index int) HistoricalForex {
		tps, _ := time.Parse("2006-01-02", item.Date)
		item.Symbol = dataBody.Symbol
		item.ID = fmt.Sprintf("%d", tps.Unix())
		item.Timestamp = tps.Unix()
		return item
	})

	return &historicalData, nil
}
