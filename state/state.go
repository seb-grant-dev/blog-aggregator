package state

import (
	"github.com/seb-grant-dev/blog-aggregator/internal/config"
)

type State struct {
	Config config.Config // This should be a pointer
}
