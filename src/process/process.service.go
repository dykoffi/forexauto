package process

import (
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

var IProcessService ProcessService

func New() *ProcessService {
	if (IProcessService != ProcessService{}) {
		return &IProcessService
	}

	loggerS := logger.New()
	dataS := data.New()

	IProcessService := ProcessService{
		logger:            loggerS,
		data:              dataS,
		fullForexQuoteDB:  "fullforexquote",
		intraDayForexDB:   "intradayforex",
		historicalForexDB: "historicalforex",
	}

	return &IProcessService

}

func (ps *ProcessService) CollectFullForexQuote() {
	fullForexQuoteData, err := ps.data.GetFullForexQuote()

	if err != nil {
		ps.logger.Error(err.Error())
		panic(err)
	}

	reqData := data.FullForexQuoteBulkData{
		Docs: fullForexQuoteData,
	}

	ioReader, err := data.TransformToReader(&reqData)

	if err != nil {
		ps.logger.Error(err.Error())
		panic(err)
	}

	if err := db.New().Insert(ps.fullForexQuoteDB, &ioReader, true); err != nil {
		return
	}
}

func (ps *ProcessService) CollectIntraDayForex() {
	intraDayForex, err := ps.data.GetIntraDayForex()

	if err != nil {
		ps.logger.Error(err.Error())
		panic(err)
	}

	reqData := data.IntraDayForexBulkData{
		Docs: intraDayForex,
	}

	ioReader, err := data.TransformToReader(&reqData)

	if err != nil {
		ps.logger.Error(err.Error())
		panic(err)
	}

	if err := db.New().Insert(ps.intraDayForexDB, &ioReader, true); err != nil {
		return
	}
}
