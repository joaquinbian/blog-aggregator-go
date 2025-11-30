package main

import (
	"blog-aggregator-go/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("error main app: %v\n", err)
		return
	}

	err = cfg.SetUser("joaquin")
	if err != nil {
		log.Fatalf("error main app: %v\n", err)
		return
	}
	fmt.Printf("configuracion: %v", cfg)
}
