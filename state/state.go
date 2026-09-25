package state

import _ "github.com/lib/pq"

import (
	"github.com/seb-grant-dev/gator/internal/config"
	"github.com/seb-grant-dev/gator/internal/database"
)

type State struct {
	DB *database.Queries
	Config *config.Config
}
