package db

import (
	"errors"
	"github.com/bneil/optimal/app/model"
	c "github.com/ostafen/clover/v2"
	"golang.org/x/exp/slog"
	"log"
)

var (
	feedCollection = "feeds"
)

func GetFeeds() (*model.BlogRoll, error) {
	d := GetInstance()
	docs, err := d.db.FindAll(c.NewQuery(feedCollection))
	if err != nil {
		slog.Error("issue in find all", err)
		return nil, err
	}
	var feeds []model.Feed
	for _, doc := range docs {
		feed := model.Feed{}
		err = doc.Unmarshal(&feed)
		if err != nil {
			slog.Warn("unable to unmarshal feed", err)
			continue
		}
		feeds = append(feeds, feed)
	}
	list := model.BlogRoll{Feeds: feeds}
	return &list, nil
}
func GetFeedById(id string) (*model.Feed, error) {
	d := GetInstance()
	log.Println("find by id", id)

	doc, err := d.db.FindFirst(c.NewQuery(feedCollection).Where(c.Field("id").Eq(id)))
	if err != nil {
		slog.Error("issue find by id", err)
		return nil, err
	}
	if doc == nil {
		return nil, errors.New("find by id failed")
	}
	feed := model.Feed{}
	err = doc.Unmarshal(&feed)
	if err != nil {
		slog.Error("couldnt unmarshal", err)
		return nil, err
	}
	return &feed, nil
}
func CreateFeed(feed *model.Feed) error {
	d := GetInstance()
	doc := c.NewDocument()
	doc.Set("id", feed.ID)
	doc.Set("text", feed.Description)
	doc.Set("title", feed.Title)
	doc.Set("type", feed.Type)
	doc.Set("html_url", feed.HtmlUrl)
	doc.Set("xml_url", feed.XMLUrl)

	log.Println("adding", feed.ID)
	output, err := d.db.InsertOne(feedCollection, doc)
	if err != nil {
		slog.Error("issue inserting", err)
		return err
	}
	slog.Info("doc:", output)
	return nil
}
func DeleteFeed(id string) bool {
	d := GetInstance()
	remQuery := c.NewQuery(feedCollection).Where(c.Field("id").Eq(id))
	err := d.db.Delete(remQuery)
	if err != nil {
		slog.Error("issue removing", err)
		return false
	}
	return true
}
func UpdateFeed(feed *model.Feed) error {
	return CreateFeed(feed)
}
