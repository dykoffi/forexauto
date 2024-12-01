package db

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockDBService struct {
	mock.Mock
}

func (m *MockDBService) Insert(database string, dataReader *io.Reader, bulk bool) error {
	args := m.Called(database, dataReader, bulk)
	return args.Error(0)
}
