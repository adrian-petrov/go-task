package main

import (
	"fmt"
	"go-task/cmd"
	"go-task/db"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
