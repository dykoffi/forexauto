package process

import (
	"time"

	"github.com/dykoffi/forexauto/src/data"
	"github.com/dykoffi/forexauto/src/db"
	"github.com/dykoffi/forexauto/src/logger"
)

type ProcessInterface interface {
	CollectFullForexQuote()
	CollectIntraDayForex()
	CollectHistoricalForex()
}

type ProcessService struct {
	fullForexQuoteDB  string
	intraDayForexDB   string
	historicalForexDB string
	logger            *logger.LoggerService
	data              *data.DataService
}

var iProcessService ProcessService

func New() *ProcessService {
	if (iProcessService != ProcessService{}) {
		return &iProcessService
	}

	iProcessService := ProcessService{
		logger:            logger.New(),
		data:              data.New(),
		fullForexQuoteDB:  "fullforexquote",
		intraDayForexDB:   "intradayforex",
		historicalForexDB: "historicalforex",
	}

	return &iProcessService

}

func (ps *ProcessService) CollectFullForexQuote() {
	ps.logger.Info("Retrieving FullForexQuote ...")
	fullForexQuoteData, err := ps.data.GetFullForexQuote()

	if err != nil {
		ps.logger.Error(err.Error())
		return
	}

	reqData := data.FullForexQuoteBulkData{
		Docs: fullForexQuoteData,
	}

	ioReader, err := data.TransformToReader(&reqData)

	if err != nil {
		ps.logger.Error(err.Error())
		return
	}

	if err := db.New().Insert(ps.fullForexQuoteDB, &ioReader, true); err != nil {
		ps.logger.Error(err.Error())
		return
	}

	ps.logger.Info("FullForexQuote saved")
}

func (ps *ProcessService) CollectIntraDayForex() {
	ps.logger.Info("Retrieving IntraDayForex ...")

	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	intraDayForex, err := ps.data.GetIntraDayForex(yesterday, yesterday)

	if err != nil {
		ps.logger.Error(err.Error())
		return
	}

	reqData := data.IntraDayForexBulkData{
		Docs: intraDayForex,
	}

	ioReader, err := data.TransformToReader(&reqData)

	if err != nil {
		ps.logger.Error(err.Error())
		return
	}

	if err := db.New().Insert(ps.intraDayForexDB, &ioReader, true); err != nil {
		ps.logger.Error(err.Error())
		return
	}

	ps.logger.Info("IntraDayForex saved")
}
