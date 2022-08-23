package cmd

import (
	"fmt"
	"go-task/db"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks the task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, len(args))
		for i, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument: %v. Please provide an id", arg)
				return
			}
			ids[i] = id
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number")
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed", id)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
