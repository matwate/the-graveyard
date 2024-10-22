package main

import (
	"errors"
	"os"

	"github.com/go-git/go-git/v5"
)

func clone_repo(link string) (repoPath string, Free func() error, err error) {

	isValid := parse_link(link)

	if !isValid {
		return "", nil, errors.New("invalid link")
	}

	dir, err := os.MkdirTemp("", "git-clone-*")
	if err != nil {
		return "", nil, err
	}

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: link,
	})
	if err != nil {
		return "", nil, err
	}

	return dir, func() error {
		return os.RemoveAll(dir)
	}, nil

}
