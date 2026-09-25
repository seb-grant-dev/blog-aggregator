package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/blog-aggregator/state"
)


func HandlerFeeds(s *state.State, cmd Command) error {

	ctx := context.Background()

	feeds, err := s.DB.GetFeedsWithUser(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s [%s] - saved by %s\n",feed.Name, feed.Url, feed.UserName)
	}
	return nil
}
