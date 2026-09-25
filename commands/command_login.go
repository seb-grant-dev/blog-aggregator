package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/gator/state"
)

func HandlerLogin(s *state.State, cmd Command) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Login expects 1 argument <username>")
	}

	username := cmd.arguments[0]

	ctx := context.Background()
	_, err := s.DB.GetUser(ctx,username)
	if err == nil {

		err := s.Config.SetUser(username)
		if err != nil {
			return err
		}
		fmt.Printf("User logged in: %s\n",username)
		return nil
	}

	return fmt.Errorf("Username [%s] not found, please try again.",username)
}
