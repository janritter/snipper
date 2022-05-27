package helper

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteToFile(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendToFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		log.Fatal(err)
	}
}
