package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var uploadS3 = &cobra.Command{
	Use:     "uploads3 [key] [filepath]",
	Short:   "upload file to s3",
	Example: `cmdvault uploads3 "commands.json" "../path" `,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		path := args[1]

		if err := awsS3client.UploadFile(context.Background(), key, path); err != nil {
			fmt.Println("❌ Error uploading:", err)
			return
		}
		fmt.Printf("✅ Saved '%s'\n", key)
	},
}

func init() {
	rootCmd.AddCommand(uploadS3)
}
