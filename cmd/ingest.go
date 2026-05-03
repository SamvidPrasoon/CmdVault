package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ingest = &cobra.Command{
	Use:     "ingest [filepath]",
	Short:   "import commands from JSON file to DB",
	Example: `cmdvault ingest "./commands.json"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		count, err := DbClient.Ingest(path)
		if err != nil {
			fmt.Println("❌ Error ingesting:", err)
			return
		}
		fmt.Printf("✅ Ingested %d commands\n", count)
	},
}

func init() {
	rootCmd.AddCommand(ingest)
}
