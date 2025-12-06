package main

import (
	"blog-aggregator-go/internal/database"
	"blog-aggregator-go/internal/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(state *State, cmd Command) error {

	name, err := utils.GetNameFromArgs(cmd.Args, cmd.Name)

	if err != nil {
		return err
	}

	_, err = state.db.GetUser(context.Background(), name)

	if err != nil {
		log.Fatal("error login: no existe el usuario")
	}

	err = state.cfg.SetUser(name)

	if err != nil {
		return err
	}

	fmt.Printf("welcome %v!\n", name)
	return nil
}

func handlerRegister(state *State, cmd Command) error {

	name, err := utils.GetNameFromArgs(cmd.Args, cmd.Name)

	if err != nil {
		return err
	}

	user, err := state.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), Name: name, CreatedAt: time.Now(), UpdatedAt: time.Now()})

	if err != nil {
		log.Fatalf("error register user: %v", err)
		return err
	}

	err = state.cfg.SetUser(name)

	if err != nil {
		return err
	}

	fmt.Printf("User %s created!\n%v", name, user)
	return nil
}
