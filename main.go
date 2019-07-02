package main

import (
	"fmt"
	"os"
	"path/filepath"
	"gtask/cmd"
	"gtask/db"
	"gtask/client"
	"gtask/domain"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, _ := homedir.Dir()
	dbPath := filepath.Join(homeDir, "gtasks.db")
	try(db.Init(dbPath))
	try(cmd.RootCmd.Execute())
}

func try(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
