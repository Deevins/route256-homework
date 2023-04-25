package cli_commands

import (
	"bufio"
	"fmt"
	"os"
)

var _ Command = (*SpellCommand)(nil)

type SpellCommand struct{}

func NewSpellCommand() *SpellCommand {
	return &SpellCommand{}
}

func (s *SpellCommand) Exec() {
	fmt.Println("Choose command you want to operate with: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Text()

	for _, char := range word {
		fmt.Print(string(char), " ")
	}
	fmt.Println()

}

func (s *SpellCommand) Info() string {
	return fmt.Sprintf("spell [word] - %s\n", "Вывести в консоль все буквы переданного на вход слова через пробел")
}
