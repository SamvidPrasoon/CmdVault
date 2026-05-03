package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samvid/cmdVault/awsS3"
	"github.com/samvid/cmdVault/config"
	"github.com/samvid/cmdVault/store"
	"github.com/spf13/cobra"
)

var DbClient *store.DbClient
var awsS3client *awsS3.S3Client
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
	cobra.OnInitialize(initDB, initS3)
}

func initDB() {
	dbPath := filepath.Join("C:/GOLANG/cmdVault/store/home/cmdvault.db")

	var err error
	DbClient, err = store.NewDbClient(dbPath)
	if err != nil {
		fmt.Println("Error opening DB:", err)
		os.Exit(1)
	} else {
		fmt.Println("DB STATUS :✅ CONNECTED")
	}
}

func initS3() {
	client, err := awsS3.NewS3Client(config.GetEnv().AWS_BUCKET, config.GetEnv().AWS_REGION)
	if err != nil {
		fmt.Println("Error connecting s3:", err)
		os.Exit(1)
	} else {
		awsS3client = client
		fmt.Println("S3 STATUS :✅ CONNECTED")
	}
}
