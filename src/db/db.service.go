package db

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/dykoffi/forexauto/src/config"
)

type DBInterface interface {
	Insert(database string, dataReader *io.Reader, bulk bool) error
}

type DBService struct {
	host     string
	username string
	password string
	client   *http.Client
}

var (
	iDBService DBService
	once       sync.Once
)

func New(config *config.ConfigService) *DBService {
	once.Do(func() {
		iDBService = DBService{
			host:     config.GetOrThrow("COUCHDB_HOST"),
			username: config.GetOrThrow("COUCHDB_USER"),
			password: config.GetOrThrow("COUCHDB_PWD"),
			client:   &http.Client{},
		}
	})

	return &iDBService

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
