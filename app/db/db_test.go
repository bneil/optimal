package db

import (
	"log"
	"testing"
)

func TestGetInstance(t *testing.T) {
	collectionName := "testing-things"
	c := GetInstance()
	err := c.db.CreateCollection(collectionName)
	if err != nil {
		t.Failed()
	}
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
		log.Println(collection)
	}
	err = c.db.DropCollection(collectionName)
	if err != nil {
		t.Failed()
	}
}
