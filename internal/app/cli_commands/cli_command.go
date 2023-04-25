package cli_commands

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
)

type Command interface {
	Exec()
	Info() string
}
