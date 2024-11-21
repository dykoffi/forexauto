package process

import (
	"sync"
	"time"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/data"
	"github.com/dykoffi/forexauto/src/db"
)

type ProcessInterface interface {
	New(config *config.ConfigService, data *data.DataService, db *db.DBService) *ProcessService
	CollectFullForexQuote() error
	CollectIntraDayForex() error
	CollectHistoricalForex() error
}

type ProcessService struct {
	fullForexQuoteDB  string
	intraDayForexDB   string
	historicalForexDB string
	config            *config.ConfigService
	data              *data.DataService
	db                *db.DBService
}

var (
	iProcessService ProcessService
	once            sync.Once
)

func New(config *config.ConfigService, data *data.DataService, db *db.DBService) *ProcessService {
	once.Do(func() {
		iProcessService = ProcessService{
			data:              data,
			db:                db,
			config:            config,
			fullForexQuoteDB:  "fullforexquote",
			intraDayForexDB:   "intradayforex",
			historicalForexDB: "historicalforex",
		}
	})

	return &iProcessService

}

func (ps *ProcessService) CollectFullForexQuote() error {
	fullForexQuoteData, err := ps.data.GetFullForexQuote()

	if err != nil {
		return err
	}

	reqData := data.FullForexQuoteBulkData{
		Docs: fullForexQuoteData,
	}

	ioReader, err := data.TransformToReader(&reqData)

	if err != nil {
		return err
	}

	if err := ps.db.Insert(ps.fullForexQuoteDB, &ioReader, true); err != nil {
		return err
	}

	return nil

}

func (ps *ProcessService) CollectIntraDayForex() error {

	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	intraDayForex, err := ps.data.GetIntraDayForex(yesterday, yesterday)

	if err != nil {
		return err
	}

	reqData := data.IntraDayForexBulkData{
		Docs: intraDayForex,
	}

	ioReader, err := data.TransformToReader(&reqData)

	if err != nil {
		return err
	}

	if err := ps.db.Insert(ps.intraDayForexDB, &ioReader, true); err != nil {
		return err
	}

	return nil

}
