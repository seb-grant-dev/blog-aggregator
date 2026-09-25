package commands

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"
)




func HandlerFollow(s *state.State, cmd command) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("This command accepts a single argument <feed url>")
	}

	currUser := s.Config.CurrentUserName

	ctx := context.Background()
	newId := uuid.New()
	feed, err := s.DB.GetFeedByUrl(ctx,cmd.arguments[0])
	if err != nil {
		return err
	}

	user, err := s.DB.GetUser(ctx,currUser)
	if err != nil {
		return err
	}

	_, err = s.DB.FollowFeed(ctx, database.FollowFeedParams{
		ID: newId,
		UserID: user.ID,
		FeedID: feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Success: %s is now following %s (%s)\n",user.Name,feed.Name,feed.Url)

	return nil


}
