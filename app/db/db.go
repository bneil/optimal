package db

import (
	"github.com/bneil/gossr_tests/config"
	c "github.com/ostafen/clover/v2"

	"sync"
)

type DB struct {
	db *c.DB
}

var instance *DB
var once sync.Once

func GetInstance() *DB {
	cfg := config.GetConfig()
	once.Do(func() {
		db, err := c.Open(cfg.Database.Location)
		if err != nil {
			panic(err)
		}
		instance = &DB{db: db}
	})
	return instance
}
