package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmds, err := DbClient.All()
		if err != nil || len(cmds) == 0 {
			fmt.Println("No commands saved yet. Use 'cmdvault add' to get started.")
			return
		}

		fmt.Printf("\n%-20s %-35s %-20s %s\n", "NAME", "COMMAND", "TAGS", "RUNS")
		fmt.Println(strings.Repeat("─", 85))

		for _, c := range cmds {
			tags := strings.Join(c.Tags, ", ")
			cmd := c.Cmd
			if len(cmd) > 33 {
				cmd = cmd[:30] + "..."
			}
			fmt.Printf("%-20s %-35s %-20s %d\n", c.Name, cmd, tags, c.RunCount)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
