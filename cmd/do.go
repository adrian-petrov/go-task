package cmd

import (
	"fmt"
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
				continue
			}
			ids[i] = id
		}
		fmt.Printf("\nThe ids are: %v", ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
