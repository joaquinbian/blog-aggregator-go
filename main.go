package main

import (
	"blog-agregator-go/internal/config"
	"fmt"
)

func main() {

	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("there was an error: %w", err)
	}

	fmt.Print(cfg.Db_url)
}
