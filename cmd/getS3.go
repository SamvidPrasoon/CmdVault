package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getS3 = &cobra.Command{
	Use:     "gets3 [key] [filepath]",
	Short:   "download file from s3",
	Example: `cmdvault gets3 "commands.json" "./downloaded.json"`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		path := args[1]

		data, err := awsS3client.GetFile(context.Background(), key)
		if err != nil {
			fmt.Println("❌ Error downloading:", err)
			return
		}

		err = os.WriteFile(path, data, 0644)
		if err != nil {
			fmt.Println("❌ Error saving file:", err)
			return
		}
		fmt.Printf("✅ Saved '%s' to '%s'\n", key, path)
	},
}

func init() {
	rootCmd.AddCommand(getS3)
}
