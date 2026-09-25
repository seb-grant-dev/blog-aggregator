package commands

import (
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"

)


func HandlerAddFeed(s *state.State, cmd command) error{
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Error: adding a feed requires a <name> and a <url>");
	}

	ctx := context.Background()

	currUser,err := s.DB.GetUser(ctx,s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	fmt.Println(currUser)

	addFeedParams := database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
		Name: cmd.arguments[0],
		Url: cmd.arguments[1],
		UserID: currUser.ID,
	}

	newFeed,err := s.DB.AddFeed(ctx,addFeedParams)
	if err != nil {
		return err
	}

	fmt.Printf("New Feed added: %+v",newFeed)

	newId := uuid.New()
	_, err = s.DB.FollowFeed(ctx, database.FollowFeedParams{
		ID: newId,
		UserID: currUser.ID,
		FeedID: newFeed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Success: %s is now following %s (%s)\n",currUser.Name,newFeed.Name,newFeed.Url)


	return nil
}
