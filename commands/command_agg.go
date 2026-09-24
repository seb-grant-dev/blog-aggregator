package commands

import (
	"fmt"
	"context"
	"html"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/feeds"
	"github.com/seb-grant-dev/blog-aggregator/models"
)

func HandlerAgg(s *state.State, cmd command) error {

	ctx := context.Background()

	feedUrl := "https://www.wagslane.dev/index.xml"
	if len(cmd.arguments) == 1 {
		feedUrl = cmd.arguments[0]
	}

	feed, err := feeds.FetchFeed(ctx,feedUrl)
	if err != nil {
		return err
	}

	cleanFeed := models.RSSFeed{}

	cleanFeed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	cleanFeed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	cleanFeed.Channel.Link = feed.Channel.Link
	cleanFeed.Channel.Item = []models.RSSItem{}


	for _, item := range feed.Channel.Item {
		cleanItem := models.RSSItem{}
		cleanItem.Title = html.UnescapeString(item.Title)
		cleanItem.Description = html.UnescapeString(item.Description)
		cleanItem.Link = item.Link
		cleanItem.PubDate = item.PubDate

		cleanFeed.Channel.Item = append(cleanFeed.Channel.Item,cleanItem)
	}

	fmt.Println(cleanFeed)

	return nil
}
