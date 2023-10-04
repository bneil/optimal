package db

import (
	"github.com/bneil/optimal/app/model"
	q "github.com/ostafen/clover/v2/query"
	"golang.org/x/exp/slog"
)

/*
A user is able to create a blog roll
A blog roll is just a collection of feeds
*/

func GetBlogrolls() (*model.BlogRoll, error) {
	d := GetInstance()

	docs, err := d.db.FindAll(q.NewQuery(feedCollection))
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
