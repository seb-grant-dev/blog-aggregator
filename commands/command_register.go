package commands

import (
	"os"
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"
)

func HandlerRegister(s *state.State, cmd Command) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Register expects 1 argument <username>")
	}

	username := cmd.arguments[0]
	ctx := context.Background()
	now := time.Now()
	userId := uuid.New()

	_, err := s.DB.GetUser(ctx,username)
	if err == nil {
		fmt.Println("A user with that name already exists")
		os.Exit(1)
	}
	
	user, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID: userId,
		CreatedAt: now,
		UpdatedAt: now,
		Name: username,
	})
	if err != nil {
		return err
	}

	s.Config.SetUser(username)

	fmt.Printf("User created: %+v\n",user)
	return nil
}
