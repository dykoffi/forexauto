package process

import (
	"errors"
	"testing"
	"time"

	"github.com/dykoffi/forexauto/src/data"
	"github.com/dykoffi/forexauto/src/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProcessTestSuite struct {
	suite.Suite
	dbService          *db.MockDBService
	dataService        *data.MockDataService
	processService     *ProcessService
	fullForexQuoteData []data.FullForexQuote
	intraForexData     []data.IntraDayForex
	date               string
}

func (suite *ProcessTestSuite) SetupSuite() {
	suite.dbService = new(db.MockDBService)
	suite.dataService = new(data.MockDataService)
	suite.processService = New(suite.dataService, suite.dbService)
	suite.fullForexQuoteData = []data.FullForexQuote{
		{ID: "", Symbol: "EURUSD", Price: 1.34563},
	}
	suite.intraForexData = []data.IntraDayForex{
		{ID: "", Open: 1.4563, Close: 1.5543, Volume: 234},
	}
	suite.date = time.Now().Add(-24 * time.Hour).Format("2006-01-02")
}

func (suite *ProcessTestSuite) TestCollectFullForexQuoteSuccess() {
	suite.dataService.On("GetFullForexQuote").Return(&suite.fullForexQuoteData, nil).Once()
	suite.dbService.On("Insert", "fullforexquote", mock.AnythingOfType("*io.Reader"), true).Return(nil).Once()
	assert.NoError(suite.T(), suite.processService.CollectFullForexQuote())

	suite.dbService.AssertExpectations(suite.T())
	suite.dataService.AssertExpectations(suite.T())
}
func (suite *ProcessTestSuite) TestCollectFullForexQuoteFailedOnDB() {
	suite.dataService.On("GetFullForexQuote").Return(&suite.fullForexQuoteData, nil).Once()
	suite.dbService.On("Insert", "fullforexquote", mock.AnythingOfType("*io.Reader"), true).Return(errors.New("Database doesn't exist")).Once()
	assert.EqualError(suite.T(), suite.processService.CollectFullForexQuote(), "Database doesn't exist")

	suite.dbService.AssertExpectations(suite.T())
	suite.dataService.AssertExpectations(suite.T())
}
func (suite *ProcessTestSuite) TestCollectFullForexQuoteFailedOnAPI() {
	suite.dataService.On("GetFullForexQuote").Return(nil, errors.New("Service not available")).Once()
	assert.EqualError(suite.T(), suite.processService.CollectFullForexQuote(), "Service not available")

	suite.dataService.AssertExpectations(suite.T())
}

func (suite *ProcessTestSuite) TestCollectIntraDayForexSuccess() {
	suite.dataService.On("GetIntraDayForex", suite.date, suite.date).Return(&suite.intraForexData, nil).Once()
	suite.dbService.On("Insert", "intradayforex", mock.AnythingOfType("*io.Reader"), true).Return(nil).Once()
	assert.NoError(suite.T(), suite.processService.CollectIntraDayForex())

	suite.dbService.AssertExpectations(suite.T())
	suite.dataService.AssertExpectations(suite.T())
}

func (suite *ProcessTestSuite) TestCollectIntraDayForexFailedOnDB() {
	suite.dataService.On("GetIntraDayForex", suite.date, suite.date).Return(&suite.intraForexData, nil).Once()
	suite.dbService.On("Insert", "intradayforex", mock.AnythingOfType("*io.Reader"), true).Return(errors.New("Database doesn't exist")).Once()
	assert.EqualError(suite.T(), suite.processService.CollectIntraDayForex(), "Database doesn't exist")

	suite.dbService.AssertExpectations(suite.T())
	suite.dataService.AssertExpectations(suite.T())
}

func (suite *ProcessTestSuite) TestCollectIntraDayForexFailedOnAPI() {
	suite.dataService.On("GetIntraDayForex", suite.date, suite.date).Return(nil, errors.New("Service not available")).Once()
	assert.EqualError(suite.T(), suite.processService.CollectIntraDayForex(), "Service not available")

	suite.dbService.AssertExpectations(suite.T())
	suite.dataService.AssertExpectations(suite.T())
}

func TestSuiteProcess(t *testing.T) {
	suite.Run(t, new(ProcessTestSuite))
}
