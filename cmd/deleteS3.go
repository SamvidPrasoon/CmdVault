package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteS3 = &cobra.Command{
	Use:     "deletes3 [key]",
	Short:   "delete file from s3",
	Example: `cmdvault deletes3 "commands.json"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		if err := awsS3client.DeleteFile(context.Background(), key); err != nil {
			fmt.Println("❌ Error deleting:", err)
			return
		}
		fmt.Printf("✅ Deleted '%s'\n", key)
	},
}

func init() {
	rootCmd.AddCommand(deleteS3)
}
