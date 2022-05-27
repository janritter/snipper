package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

func getGitUrlParts(repo string) (string, string) {
	parts := strings.Split(repo, ":")
	if len(parts) == 1 {
		fmt.Println("Repo is not in the format provider:user/repo")
		os.Exit(1)
	}
	provider := parts[0]

	switch provider {
	case "gh":
		return "git@github.com:" + parts[1] + ".git", parts[1]
	default:
		fmt.Println("Provider " + provider + " not supported")
		os.Exit(1)
	}

	return "", ""
}

func checkOrCloneRepo(gitUrl string, path string) {
	_, err := git.PlainOpen(path)
	if err != nil {
		fmt.Println("Snipper collection locally not found, cloning from " + gitUrl)
		_, err = git.PlainClone(path, false, &git.CloneOptions{
			URL:      gitUrl,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func updateRepo(gitUrl string, path string) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Local copy found, updating...")
	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatal(err)
	}
}
