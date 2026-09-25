package commands

import (
	"fmt"
	"context"
	"github.com/seb-grant-dev/blog-aggregator/state"
)

func HandlerReset(s *state.State, cmd Command) error {
	ctx := context.Background()
	err := s.DB.Reset(ctx)
	if err != nil {
		return fmt.Errorf("Error resetting database: %s",err)
	}

	return nil
}
