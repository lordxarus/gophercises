/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"gophercises/task/cmd"
	"gophercises/task/db"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

/*
TODO: Only update our "view" of the tasks once the user lists them again

Ex:

  - Add two tasks
    1. Do A
    2. Do B

  - Run "task do 1"

  - Run "task do 2"

  - Program will error because "Do B" is now at 1

  - If we stored the list we get from db.AllTasks()
    and only updated it when the user runs "task list"
    again we would have behavior like:

  - Add two tasks

  - "task do 1"

  - "task do 2"

  - Completes both tasks
*/
func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
