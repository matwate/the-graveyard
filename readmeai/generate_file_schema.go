package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func generate_repo_schema(repo_path string) (err error) {

	err = filepath.WalkDir(repo_path, file_schema_generator)
	if err != nil {
		return err
	}
	err = nil
	return

}

func file_schema_generator(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.IsDir() {
		return nil
	}
	// Create a new file schema
	if strings.Contains(path, "\\.git\\") {
		return nil
	}
	NewFileSchema(path)
	return nil
}
