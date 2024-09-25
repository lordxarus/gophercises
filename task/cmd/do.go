/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		fmt.Printf("%+v\n", ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
