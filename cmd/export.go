package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [filename]",
	Short: "Export all commands to a JSON file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmds, err := DbClient.All()
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		data, _ := json.MarshalIndent(cmds, "", "  ")
		os.WriteFile(args[0], data, 0644)
		fmt.Printf("✅ Exported %d commands to %s\n", len(cmds), args[0])
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
