package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/janritter/snipper/helper"
	"github.com/spf13/cobra"
)

var outputFilename string
var appendFilename string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [collection url] [snippet]",
	Short: "Tool to get various snippets directly from your CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 2 {
			fmt.Println("At least two args (collection url and snippet) are required, multiple snippet args for nested collections are supported")
			fmt.Println("Example: snipper get gh:username/snipper-collection terraform state s3")
			os.Exit(1)
		}

		gitUrl, repo := getGitUrlParts(args[0])
		path := "/tmp/snipper/" + repo

		checkOrCloneRepo(gitUrl, path)

		path = path + "/" + strings.Join(args[1:], "/")
		filePath := path + ".md"

		// Check if file exists
		if _, err := os.Stat(filePath); err == nil {
			renderOutputAppendSnippet(filePath)
		}

		// Check if directory
		if fi, err := os.Stat(path); err == nil {
			if fi.IsDir() {
				files, err := ioutil.ReadDir(path)
				if err != nil {
					log.Fatal(err)
				}

				filenames := []string{}
				for _, f := range files {
					if !f.IsDir() {
						filenames = append(filenames, f.Name())
					}
				}

				if len(filenames) == 0 {
					log.Fatal("No snippets found in directory: " + path)
				}

				if len(filenames) == 1 {
					fmt.Println("Only one snippet found in directory, outputting...")
					renderOutputAppendSnippet(path + "/" + filenames[0])
				}

				choice := helper.GetList(filenames)

				if choice == "" {
					os.Exit(0)
				}

				renderOutputAppendSnippet(path + "/" + choice)

				os.Exit(0)
			}
		}

		fmt.Println("No file or directory for values found: " + path)
		os.Exit(1)
	},
}

func init() {
	getCmd.Flags().StringVarP(&outputFilename, "output", "o", "", "Output snippet to file (overwrites existing file)")
	getCmd.Flags().StringVarP(&appendFilename, "append", "a", "", "Append snippet to file")

	rootCmd.AddCommand(getCmd)
}
