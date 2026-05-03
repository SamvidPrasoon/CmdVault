package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/samvid/cmdVault/store"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [name] [command]",
	Short:   "Save a new command",
	Example: `  cmdvault add deploy-prod "kubectl apply -f ./k8s/" --desc "Deploy to prod" --tags k8s,infra`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		command := args[1]
		desc, _ := cmd.Flags().GetString("desc")
		tagsStr, _ := cmd.Flags().GetString("tags")

		tags := []string{}
		if tagsStr != "" {
			tags = strings.Split(tagsStr, ",")
		}

		c := store.Command{
			ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
			Name:        name,
			Cmd:         command,
			Description: desc,
			Tags:        tags,
			CreatedAt:   time.Now(),
		}

		if err := DbClient.Save(c); err != nil {
			fmt.Println("❌ Error saving:", err)
			return
		}
		fmt.Printf("✅ Saved '%s'\n", name)
	},
}

func init() {
	addCmd.Flags().String("desc", "", "Description of the command")
	addCmd.Flags().String("tags", "", "Comma-separated tags (e.g. k8s,infra,docker)")
	rootCmd.AddCommand(addCmd)
}
