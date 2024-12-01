package data

import "github.com/stretchr/testify/mock"

type MockDataService struct {
	mock.Mock
}

func (m *MockDataService) GetFullForexQuote() (*[]FullForexQuote, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*[]FullForexQuote), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataService) GetIntraDayForex(from string, to string) (*[]IntraDayForex, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*[]IntraDayForex), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataService) GetHistoricalDailyForex() (*[]HistoricalForex, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*[]HistoricalForex), args.Error(1)
	}
	return nil, args.Error(1)
}
