package main

import (
	"fmt"
	"github.com/seb-grant-dev/blog-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	fmt.Printf("Config: %+v\n",cfg)

	fmt.Println("Update username to seb-grant-dev")
	cfg.SetUser("seb-grant-dev")


	
	cfgUpdated := config.Read()
	fmt.Printf("Config: %+v\n",cfgUpdated)
}
