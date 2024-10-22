package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOllamaEnv(t *testing.T) {
	err := load_ollama_env()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestFillSchemaWithLlm(t *testing.T) {
	/*repo_link := "https://github.com/matwate/golgc.git"
	repo_path, Free, err := clone_repo(repo_link)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	defer Free()
	err = generate_repo_schema(repo_path)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	fsch := global_schemas[0]
	err = fill_schema_with_llm(fsch)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	fsch.Print()
	*/
}

/*func TestFullRepoSchemaGeneration(t *testing.T) {
	repo_link := "https://github.com/matwate/golgc.git"
	repo_path, _, err := clone_repo(repo_link)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = generate_repo_schema(repo_path)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	for _, fsch := range global_schemas {
		err = fill_schema_with_llm(fsch)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	}
	//save_results()

}*/

func save_results() {
	// Create the Summaries directory if it doesn't exist

	err := os.MkdirAll("./Summaries", os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Iterate over the schemas array and write each summary to a file
	for i, fsch := range global_schemas {
		filePath := filepath.Join("./Summaries", fmt.Sprintf("summary_%d.txt", i))
		err := os.WriteFile(filePath, []byte(fsch.summary), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		}
	}

}
