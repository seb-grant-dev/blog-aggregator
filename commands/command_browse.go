package commands

import (
	"fmt"
	"context"
	"strconv"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"

)

func HandlerBrowse(s *state.State, cmd Command, currUser database.User) error {

	limit := 2
	if len(cmd.arguments) > 0 {
		intLimit,err := strconv.Atoi(cmd.arguments[0])
		if err == nil {
			limit = intLimit
		}

	}

	ctx := context.Background()

	postsParams := database.GetPostsForUserParams{
		Name: currUser.Name,
		Limit: int32(limit),
	}
	posts, err := s.DB.GetPostsForUser(ctx, postsParams)



	if err != nil {
		return err
	}

	for p,post := range posts {
		fmt.Printf("%d. %s\n   %s\n", p + 1, post.Title, post.Url)
	}

	return nil
}
