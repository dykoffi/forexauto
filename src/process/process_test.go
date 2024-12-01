package process

import (
	"errors"
	"testing"

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
}

func (suite *ProcessTestSuite) SetupSuite() {
	suite.dbService = new(db.MockDBService)
	suite.dataService = new(data.MockDataService)
	suite.processService = New(suite.dataService, suite.dbService)
	suite.fullForexQuoteData = []data.FullForexQuote{
		{ID: "", Symbol: "EURUSD", Price: 1.34563},
	}
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
	assert.Error(suite.T(), suite.processService.CollectFullForexQuote())

	suite.dbService.AssertExpectations(suite.T())
	suite.dataService.AssertExpectations(suite.T())
}
func (suite *ProcessTestSuite) TestCollectFullForexQuoteFailedOnAPI() {
	suite.dataService.On("GetFullForexQuote").Return(nil, errors.New("Service not available")).Once()
	assert.Error(suite.T(), suite.processService.CollectFullForexQuote())

	suite.dataService.AssertExpectations(suite.T())
}

func TestSuiteProcess(t *testing.T) {
	suite.Run(t, new(ProcessTestSuite))
}
