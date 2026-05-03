package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a saved command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := DbClient.Delete(args[0]); err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
		fmt.Printf("🗑️  Deleted '%s'\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
