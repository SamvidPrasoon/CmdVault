package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Execute a saved command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := DbClient.GetCmd(args[0])
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		fmt.Printf("▶ Running: %s\n\n", c.Cmd)

		// Execute the command in shell
		shell := exec.Command("sh", "-c", c.Cmd)
		shell.Stdout = os.Stdout
		shell.Stderr = os.Stderr
		shell.Stdin = os.Stdin

		if err := shell.Run(); err != nil {
			fmt.Println("❌ Command failed:", err)
			return
		}

		// Track how many times it was run
		DbClient.IncrementRunCount(c.Name)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
