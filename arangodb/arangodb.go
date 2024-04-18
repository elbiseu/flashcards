package arangodb

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type Connection struct{}

type ArangoDB struct {
	Client    driver.Client
	endpoints []string
	password  string
	username  string
}

func NewArangoDB(username, password string, endpoints []string) (*ArangoDB, error) {
	arangoDB := &ArangoDB{
		endpoints: endpoints,
		username:  username,
		password:  password,
	}
	if err := arangoDB.connect(); err != nil {
		return nil, err
	}
	return arangoDB, nil
}

func (a *ArangoDB) connect() error {
	connection, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: a.endpoints,
	})
	if err != nil {
		return err
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     connection,
		Authentication: driver.BasicAuthentication(a.username, a.password),
	})
	if err != nil {
		return err
	}
	a.Client = client
	return nil
}
