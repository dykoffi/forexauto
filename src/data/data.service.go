package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/logger"
)

type DataInterface interface {
	New() *DataService
	GetFullForexQuote() *FullForexQuote
	GetIntraDayForex() *IntraDayForex
	GetDailyForex() *DailyForex
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
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return nil, err
	}

	var dataBody []FullForexQuote

	if err := json.Unmarshal(body, &dataBody); err != nil {
		return nil, err
	}

	return &dataBody, nil

}

func (ds *DataService) GetIntraDayForex() *IntraDayForex {
	return &IntraDayForex{}
}

func (ds *DataService) GetDailyForex() *DailyForex {
	return &DailyForex{}
}
