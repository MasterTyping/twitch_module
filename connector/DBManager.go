package connector

import (
	"database/sql/driver"
	"sync"
)

type dbmanager struct {
	list map[string]driver.Connector
}

var lock = &sync.Mutex{}

var DBManager *dbmanager

func (db *dbmanager) GetConnection() {
	db.list = make(map[string]driver.Connector)

}

func Init() (*dbmanager, error) {
	if DBManager == nil {
		lock.Lock()
		defer lock.Unlock()
		if DBManager == nil {
			DBManager = &dbmanager{}
		}
	}
	return DBManager, nil
}
