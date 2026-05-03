package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samvid/cmdVault/store"
	"github.com/spf13/cobra"
)

var DbClient *store.DbClient
var rootCmd = &cobra.Command{
	Use:   "cmdvault",
	Short: "🗄️  cmdvault — Save and run your infra/dev commands",
	Long: `
cmdvault is a CLI tool for developers to save, organize,
search and run frequently used shell/infra commands.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initDB)
}

func initDB() {
	dbPath := filepath.Join("C:/GOLANG/cmdVault/store/home/cmdvault.db")

	var err error
	DbClient, err = store.NewDbClient(dbPath)
	if err != nil {
		fmt.Println("Error opening DB:", err)
		os.Exit(1)
	}
}
