package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"

	"github.com/janritter/snipper/helper"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var outputFilename string
var appendFilename string

func renderOutputAppendSnippet(filepath string) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	out, err := glamour.Render(string(content), "dark")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)

	if outputFilename != "" {
		codeblock := helper.GetCodeblockFromMarkdown(string(content))

		helper.WriteToFile(outputFilename, codeblock)
	}

	if appendFilename != "" {
		codeblock := helper.GetCodeblockFromMarkdown(string(content))

		helper.AppendToFile(appendFilename, codeblock)
	}
	os.Exit(0)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snipper",
	Short: "Tool to get various snippets directly from your CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 2 {
			fmt.Println("At least two args (provider url and snippet) are required, multiple snippet args for nested collections are supported")
			fmt.Println("Example: snipper gh:username/snipper-collection terraform state s3")
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.snipper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&outputFilename, "output", "o", "", "Output snippet to file (overwrites existing file)")
	rootCmd.Flags().StringVarP(&appendFilename, "append", "a", "", "Append snippet to file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".snipper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".snipper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
