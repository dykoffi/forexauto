package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/logger"
	lop "github.com/samber/lo/parallel"
)

type DataInterface interface {
	New() *DataService
	GetFullForexQuote() *FullForexQuote
	GetIntraDayForex() *IntraDayForex
	GetDailyForex() *HistoricalDailyForex
}

type DataService struct {
	host   string
	apiKey string
	logger *logger.LoggerService
}

var IDataService DataService

func New() *DataService {

	if (IDataService != DataService{}) {
		return &IDataService
	}

	config := config.New()
	logger := logger.New()

	IDataService = DataService{
		apiKey: config.GetOrThrow("FOREX_API_KEY"),
		host:   config.GetOrThrow("FOREX_BASE_URL"),
		logger: logger,
	}

	return &IDataService
}

func (ds *DataService) GetFullForexQuote(symbol string) (*[]FullForexQuote, error) {

	path := fmt.Sprintf("/quote/%s", symbol)
	finalUrl := fmt.Sprintf("%s/%s?apikey=%s", ds.host, path, ds.apiKey)
	res, err := http.Get(finalUrl)

	if err != nil {
		ds.logger.Error(err.Error())
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ds.logger.Error(err.Error())
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return nil, err
	}

	var dataBody []FullForexQuote

	if err := json.Unmarshal(body, &dataBody); err != nil {
		ds.logger.Error(err.Error())
		return nil, err
	}

	return &dataBody, nil
}

func (ds *DataService) GetIntraDayForex() *IntraDayForex {
	return &IntraDayForex{}
}

func (ds *DataService) GetHistoricalDailyForex(symbol string) (*[]HistoricalForex, error) {

	path := fmt.Sprintf("/historical-price-full/%s", symbol)
	finalUrl := fmt.Sprintf("%s/%s?apikey=%s", ds.host, path, ds.apiKey)
	res, err := http.Get(finalUrl)

	if err != nil {
		ds.logger.Error(err.Error())
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ds.logger.Error(err.Error())
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return nil, err
	}

	var dataBody HistoricalDailyForex

	if err := json.Unmarshal(body, &dataBody); err != nil {
		ds.logger.Error(err.Error())
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
