package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search saved commands by name, tag or description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, err := DbClient.Search(args[0])
		if err != nil || len(results) == 0 {
			fmt.Printf("No results for '%s'\n", args[0])
			return
		}

		fmt.Printf("\n🔍 Results for '%s':\n\n", args[0])
		for _, c := range results {
			fmt.Printf("  %-20s %s\n", c.Name, c.Cmd)
			if c.Description != "" {
				fmt.Printf("  %-20s %s\n", "", "|-- "+c.Description)
			}
			if len(c.Tags) > 0 {
				fmt.Printf("  %-20s Tags: %s\n", "", strings.Join(c.Tags, ", "))
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
