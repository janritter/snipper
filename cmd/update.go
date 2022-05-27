package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use: "update [collection url]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 1 {
			fmt.Println("At least one arg (collection url) is required")
			fmt.Println("Example: snipper update gh:username/snipper-collection")
			os.Exit(1)
		}

		gitUrl, repo := getGitUrlParts(args[0])
		path := "/tmp/snipper/" + repo

		updateRepo(gitUrl, path)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
