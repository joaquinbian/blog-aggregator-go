package main

import (
	"blog-aggregator-go/internal/config"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type State struct {
	cfg *config.Config
}

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalf("expected 2 or more arguements")
	}

	cfgFile, err := config.Read()

	if err != nil {
		log.Fatalf("error main app: %v\n", err)
		return
	}

	state := &State{
		cfg: &cfgFile,
	}

	cmds := Commands{
		commands: make(map[string]func(*State, Command) error),
	}

	cmds.Register("login", handlerLogin)

	var cmdName = args[1]
	var cmdArgs = args[2:]
	err = cmds.Run(state, Command{Name: cmdName, Args: cmdArgs})

	if err != nil {
		log.Fatalf("error main app: %v\n", err)
		return
	}
}
