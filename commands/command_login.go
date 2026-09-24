package commands

import (
	"fmt"
	"github.com/seb-grant-dev/blog-aggregator/state"
)

func HandlerLogin(s *state.State, cmd command) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Login expects 1 argument <username>")
	}

	username := cmd.arguments[0]
	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User logged in: %s\n",username)
	return nil
}
