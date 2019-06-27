package commands

import (
	Logger "github.com/sirupsen/logrus"
)

//Command ...
type Command interface {
	Execute(Command) error
}

//type AddCommand func(db *bolt.DB, logger)

func addCommand(db *bolt.DB, log Logger) func(Command) error {
	return func(command Command) {
		log.init()
	}
}

func (addCommand *AddCommand) execute() error {

}

// type Command func()

// func AddCommand(command Command) error {

// }

// func DoneCommand(command Command) error {

// }

// func EditCommand(command Command) error {

// }
// func ListCommand(command Command) error {

// }
// func LoginCommand(command Command) error {

// }
// func PullCommand(command Command) error {

// }
// func PushCommand(command Command) error {

// }
// func RemoveCommand(command Command) error {

// }
// func VersionCommand(command Command) error {

// }
