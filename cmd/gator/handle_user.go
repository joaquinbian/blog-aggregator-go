package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joaquinbian/blog-aggregator-go/internal/database"
	"github.com/joaquinbian/blog-aggregator-go/internal/utils"
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

	fmt.Printf("User %s created!\n%v\n", name, user)
	return nil
}

func handlerList(state *State, cmd Command) error {
	users, err := state.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("error al obtener usuarios: %v", err)
	}

	for _, u := range users {
		if u.Name == state.cfg.Current_user_name {
			fmt.Printf("* %v (current)\n", u.Name)
		} else {
			fmt.Printf("* %v\n", u.Name)
		}
	}
	return nil
}
