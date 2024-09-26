/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do a task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong fetching tasks:", err)
		}

		for _, id := range ids {
			// our tasks start at idx 1
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number:", id)
				continue
			}
			// convert to zero based idx
			task := tasks[id-1]
			err = db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Error: %s occured while marking \"%d\" as complete\n", err, id)
			} else {
				fmt.Printf("Marked \"%d\" as complete\n", id)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
