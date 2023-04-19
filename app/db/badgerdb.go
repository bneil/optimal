package db

import (
	"github.com/bneil/gossr_tests/config"
	c "github.com/ostafen/clover/v2"
	"golang.org/x/exp/slog"
	"log"

	"sync"
)

type DB struct {
	db   *c.DB
	Lock sync.Mutex
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

const FEED_COLLECTION = "feeds"

func (d *DB) SetupDb() error {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	hasFeeds, err := d.db.HasCollection(FEED_COLLECTION)
	if err != nil {
		log.Println("couldnt read the db")
		return err
	}
	if !hasFeeds {
		err := d.db.CreateCollection(FEED_COLLECTION)
		if err != nil {
			return err
		}
	} else {
		slog.Info("already have a", FEED_COLLECTION)
	}
	return nil
}
