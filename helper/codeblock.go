package helper

import (
	"strings"
)

func GetCodeblockFromMarkdown(markdown string) string {
	var codeblock string

	lines := strings.Split(markdown, "\n")
	foundCodeblock := false
	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			foundCodeblock = !foundCodeblock
		} else {
			if foundCodeblock {
				codeblock += line + "\n"
			}
		}
	}

	return codeblock
}
