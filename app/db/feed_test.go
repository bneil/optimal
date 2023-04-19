package db

import (
	"fmt"
	"github.com/bneil/gossr_tests/app/model"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"testing"
)

var (
	testCollection = "TEST_COLLECTION"
)

func newFeed(id string) *model.Feed {
	if id == "" {
		id = uuid.New().String()
	}
	return &model.Feed{
		ID:      id,
		Title:   fmt.Sprintf("Title:%d", rand.Int()),
		Text:    fmt.Sprintf("Text:%d", rand.Int()),
		Type:    "temp",
		HtmlUrl: "http://testing.com",
		XMLUrl:  "http://xml.com",
	}
}

func TestGetInstance(t *testing.T) {
	setup()
	collectionName := testCollection
	c := GetInstance()
	hasCollection, err := c.db.HasCollection(collectionName)
	if err != nil {
		t.Failed()
	}
	if hasCollection {
		log.Println("collection was created")
	}
	collections, err := c.db.ListCollections()
	if err != nil {
		return
	}
	for _, collection := range collections {
		log.Println("collection:", collection)
	}
	shutdown()
}

func Test_GetFeeds(t *testing.T) {
	setup()
	err := CreateFeed(newFeed(""))
	if err != nil {
		t.Fatalf("got %v", err)
	}
	lst, err := GetFeeds()
	if err != nil {
		t.Fatalf("got %v", err)
	}
	for i, i2 := range lst.Feeds {
		log.Println(i, ":", i2)
	}
	shutdown()
}
func Test_GetFeedById(t *testing.T) {
	setup()
	sampleId := "test:12345"
	err := CreateFeed(newFeed(sampleId))
	if err != nil {
		t.Fatalf("got %v", err)
	}

	log.Println("finding")
	doc, err := GetFeedById(sampleId)
	if err != nil {
		t.Fatalf("got %v", err)
	}
	log.Println(doc)
	shutdown()
}

func setup() {
	feedCollection = testCollection
	c := GetInstance()
	hasCollection, err := c.db.HasCollection(testCollection)
	if err != nil {
		panic(err)
	}
	if !hasCollection {
		err := c.db.CreateCollection(testCollection)
		if err != nil {
			panic(err)
		}
	}
}

func shutdown() {
	c := GetInstance()
	err := c.db.DropCollection(testCollection)
	if err != nil {
		log.Println("huh?", err)
		panic(err)
	}
}
