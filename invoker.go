package main

import (
	"gtask/main"
)
//Invoker for command pattern...
type Invoker struct {
}

func (invoker *Invoker) executeCommands(command Command) error {
	command.Execute()
}
