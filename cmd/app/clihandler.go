package main

import (
	"bufio"
	"fmt"
	"os"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/cli_commands"
)

const helpCommand = "help"

type CLIHandler struct {
	commands map[string]cli_commands.Command
}

func NewCLIHandler() *CLIHandler {
	return &CLIHandler{
		commands: make(map[string]cli_commands.Command),
	}
}

func (s *CLIHandler) Register(alias string, command cli_commands.Command) {
	if alias == "help" {
		fmt.Println("You can not register help command")
		return
	}

	s.commands[alias] = command
}

func (s *CLIHandler) CommandProcessing() {
	fmt.Println("Type your command:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()
	if command == helpCommand {
		s.help()
		return
	}

	if v, ok := s.commands[command]; ok {
		v.Exec()
	} else {
		fmt.Println("Unknown command")
	}
}

func (s *CLIHandler) help() {
	fmt.Println("List of commands:")
	for k, v := range s.commands {
		fmt.Printf("%s - %s\n", k, v.Info())
	}
}
