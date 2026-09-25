package commands

import (
	"fmt"
	"time"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/feeds"
)

func HandlerAgg(s *state.State, cmd Command) error {

	// ctx := context.Background()

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Error: This command requires a duration, such as 30s, 2m, 1h etc")
	}

	time_between_reqs,err := time.ParseDuration(cmd.arguments[0])
	fmt.Printf("Collecting feeds every %s\n",time_between_reqs)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <- ticker.C {
		feeds.ScrapeFeeds(s)
	}

	return nil


/*

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

	*/
}
