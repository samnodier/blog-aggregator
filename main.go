package main

import (
	"fmt"
	"github.com/samnodier/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	err = cfg.SetUser("sam")
	if err != nil {
		fmt.Println(err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
