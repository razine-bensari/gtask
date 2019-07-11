package main

import (
	"gtask/cmd"
	"gtask/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, _ := homedir.Dir()
	dbPath := filepath.Join(homeDir, "gtasks.db")
	db.Init(dbPath)
	cmd.RootCmd.Execute()

}
