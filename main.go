package main

import (
	"fmt"
	"os"
	"sync"

	"gtask/cmd"
	"gtask/oauth"
)

func main() {

	try(cmd.RootCmd.Execute())
}

func try(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
