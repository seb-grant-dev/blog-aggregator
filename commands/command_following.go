package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/gator/state"
	"github.com/seb-grant-dev/gator/internal/database"
)




func HandlerFollowing(s *state.State, cmd Command, currUser database.User) error {

	if len(cmd.arguments) != 0 {
		fmt.Errorf("This command does not require any inputs. Ignoring...")
	}

	ctx := context.Background()

	feeds, err := s.DB.GetFeedFollowsForUser(ctx,currUser.Name)
	if err != nil {
		return err
	}

	fmt.Println(feeds)


	fmt.Printf("%s is following:\n",currUser.Name)
	for _, feed := range feeds {
		fmt.Printf(" - %s\n", feed.FeedName)
	}

	return nil


}
