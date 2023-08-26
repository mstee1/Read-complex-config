package main

import (
	"fmt"

	"github.com/mstee1/Read-complex-config/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	if cfg.Version != "" {
		fmt.Println(cfg.Version)
		return
	}
	for _, v := range cfg.Workers {
		fmt.Println(v)
	}
}
