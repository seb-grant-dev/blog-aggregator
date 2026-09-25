
package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/gator/state"
	"github.com/seb-grant-dev/gator/internal/database"
)




func HandlerUnfollow(s *state.State, cmd Command, currUser database.User) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("This command accepts a single argument <feed url>")
	}

	ctx := context.Background()
	feed, err := s.DB.GetFeedByUrl(ctx,cmd.arguments[0])
	if err != nil {
		return err
	}

	err = s.DB.UnfollowFeed(ctx, database.UnfollowFeedParams{
		UserID: currUser.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Success: %s is no longer following %s (%s)\n",currUser.Name,feed.Name,feed.Url)

	return nil


}
