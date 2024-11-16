package db

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/logger"
)

type DBInterface interface {
}

type DBService struct {
	host          string
	username      string
	password      string
	databaseRead  string
	databaseWrite string
	logger        *logger.LoggerService
	client        *http.Client
}

var IDBService DBService

func New() *DBService {
	if (IDBService != DBService{}) {
		return &IDBService
	}

	config := config.New()
	logger := logger.New()

	IDBService = DBService{
		host:          config.GetOrThrow("COUCHDB_HOST"),
		username:      config.GetOrThrow("COUCHDB_USER"),
		password:      config.GetOrThrow("COUCHDB_PWD"),
		databaseRead:  config.GetOrThrow("COUCHDB_DB_READ"),
		databaseWrite: config.GetOrThrow("COUCHDB_DB_WRITE"),
		logger:        logger,
		client:        &http.Client{},
	}

	return &IDBService

}

func (dbs *DBService) AddData(data io.Reader) {
	fullPath := fmt.Sprintf("%s/%s/_bulk_docs", dbs.host, dbs.databaseWrite)
	req, _ := http.NewRequest(http.MethodPost, fullPath, data)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Length", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.SetBasicAuth(dbs.username, dbs.password)

	_, err := dbs.client.Do(req)

	if err != nil {
		panic(err)
	}

	req.Body.Close()

}
