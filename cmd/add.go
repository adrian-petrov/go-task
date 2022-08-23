package cmd

import (
	"fmt"
	"go-task/db"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong when adding the task")
			return
		}
		fmt.Printf("Added task: \"%s\"", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
