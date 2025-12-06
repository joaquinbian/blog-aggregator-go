package main

import (
	"blog-aggregator-go/internal/config"
	"blog-aggregator-go/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	//chequeamos argumentos necesarios
	args := os.Args

	if len(args) < 2 {
		log.Fatalf("expected 2 or more arguements")
	}

	//leemos el archivo de configuracion
	cfgFile, err := config.Read()

	if err != nil {
		log.Fatalf("error main app: %v\n", err)
		return
	}

	//conectamos a la db
	db, err := sql.Open("postgres", cfgFile.Db_url)

	if err != nil {
		log.Fatalf("error main app: %v", err)
	}

	//agregamos la conexion para las queries al estado
	dbQueries := database.New(db)

	state := &State{
		cfg: &cfgFile,
		db:  dbQueries,
	}

	//inicializamos comandos
	cmds := Commands{
		commands: make(map[string]func(*State, Command) error),
	}

	//registramos comandos
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)

	//corremos el comando ingresado por el usuario
	var cmdName = args[1]
	var cmdArgs = args[2:]
	err = cmds.Run(state, Command{Name: cmdName, Args: cmdArgs})

	if err != nil {
		log.Fatalf("error main app: %v\n", err)
		return
	}
}
