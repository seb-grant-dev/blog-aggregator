
package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/blog-aggregator/state"
)

func HandlerUsers(s *state.State, cmd Command) error {

	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Get Users expects 0 arguments")
	}

	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx)
	if err == nil {

		currUser := s.Config.CurrentUserName

		for _, user := range users {
			if user.Name == currUser {
				fmt.Printf("* %s (current)\n",user.Name)
			} else {
				fmt.Printf("* %s\n",user.Name)
			}
		}

		return nil
	}

	return fmt.Errorf("Error retrieving users, please try again.")
}
