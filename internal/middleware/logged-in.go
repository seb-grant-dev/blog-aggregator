package middleware

import (
	"context"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"
	"github.com/seb-grant-dev/blog-aggregator/commands"
)

func MiddlewareLoggedIn(handler func(s *state.State, cmd commands.Command, user database.User) error) func (*state.State, commands.Command) error {

	return func(s *state.State, cmd commands.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s,cmd,user)
	}


}
