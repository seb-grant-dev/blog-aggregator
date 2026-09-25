package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/blog-aggregator/state"
)




func HandlerFollowing(s *state.State, cmd command) error {

	if len(cmd.arguments) != 0 {
		fmt.Errorf("This command does not require any inputs. Ignoring...")
	}

	currUser := s.Config.CurrentUserName

	ctx := context.Background()

	feeds, err := s.DB.GetFeedFollowsForUser(ctx,currUser)
	if err != nil {
		return err
	}

	fmt.Println(feeds)


	fmt.Printf("%s is following:\n",currUser)
	for _, feed := range feeds {
		fmt.Printf(" - %s\n", feed.FeedName)
	}

	return nil


}
