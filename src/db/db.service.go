package db

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/logger"
)

type DBInterface interface {
	New() *DBService
}

type DBService struct {
	host     string
	username string
	password string
	logger   *logger.LoggerService
	client   *http.Client
}

var IDBService DBService

func New() *DBService {
	if (IDBService != DBService{}) {
		return &IDBService
	}

	config := config.New()

	IDBService = DBService{
		host:     config.GetOrThrow("COUCHDB_HOST"),
		username: config.GetOrThrow("COUCHDB_USER"),
		password: config.GetOrThrow("COUCHDB_PWD"),
		logger:   logger.New(),
		client:   &http.Client{},
	}

	return &IDBService

}

func (dbs *DBService) Insert(database string, dataReader *io.Reader, bulk bool) error {

	var fullPath string

	if bulk {
		fullPath = fmt.Sprintf("%s/%s/_bulk_docs", dbs.host, database)
	} else {
		fullPath = fmt.Sprintf("%s/%s", dbs.host, database)
	}

	req, err := http.NewRequest(http.MethodPost, fullPath, *dataReader)

	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	// req.Header.Add("Content-Length", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.SetBasicAuth(dbs.username, dbs.password)

	if _, err := dbs.client.Do(req); err != nil {
		return err
	}

	req.Body.Close()

	return nil
}
