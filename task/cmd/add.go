package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task.")
			return
		}
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Printf("something went wrong: %v", err)
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
