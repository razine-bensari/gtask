package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gtask/cmd"
	"gtask/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	try(db.Init(dbPath))
	try(cmd.RootCmd.Execute())
}

func try(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
