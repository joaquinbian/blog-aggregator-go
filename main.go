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

	cfg.SetUser("Joaquin")

	fmt.Print(cfg)

}
