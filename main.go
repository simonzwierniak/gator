package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error: connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	appState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	appCommands := commands{
		registry: make(map[string]func(*state, command) error),
	}
	appCommands.register("login", handlerLogin)
	appCommands.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	err = appCommands.run(appState, command{name: commandName, args: commandArgs})
	if err != nil {
		log.Fatal(err)
	}
}
