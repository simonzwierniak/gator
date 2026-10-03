package main

import (
	"gator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	appState := state{cfg: &cfg}
	appCommands := commands{
		registry: make(map[string]func(*state, command) error),
	}
	appCommands.register("login", handlerLogin)
	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("insert a command")
		os.Exit(1)	
	}
	userCommand := command{name: userArgs[1], args: userArgs[2:]}
	err = appCommands.run(&appState, userCommand)
	if err != nil {
		fmt.Println("error in the process")
		os.Exit(1)
	}
}