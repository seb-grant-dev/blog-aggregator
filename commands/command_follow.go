package commands

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/seb-grant-dev/gator/state"
	"github.com/seb-grant-dev/gator/internal/database"
)




func HandlerFollow(s *state.State, cmd Command, currUser database.User) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("This command accepts a single argument <feed url>")
	}


	ctx := context.Background()
	newId := uuid.New()
	feed, err := s.DB.GetFeedByUrl(ctx,cmd.arguments[0])
	if err != nil {
		return err
	}

	_, err = s.DB.FollowFeed(ctx, database.FollowFeedParams{
		ID: newId,
		UserID: currUser.ID,
		FeedID: feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Success: %s is now following %s (%s)\n",currUser.Name,feed.Name,feed.Url)

	return nil


}
