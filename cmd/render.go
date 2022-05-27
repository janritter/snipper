package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/janritter/snipper/helper"
)

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
