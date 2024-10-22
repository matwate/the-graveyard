package main

import "testing"

func TestReadmeGenera(t *testing.T) {
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
	prompt_llm_model()
}
