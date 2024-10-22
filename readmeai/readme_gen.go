package main

import (
	"fmt"
	"os"
	"strings"
)

// It will take all of the global schemas summaries and build the readme on it

func compose_summaries_prompt() string {
	header := `You are tasked with creating a README file in markdown format for a repository. Here's a curated list of summaries for the fields in the repositories:
	- Use emojis to make your README more visually appealing and engaging.
	- Consider adding anchors to sections of your README to allow easy navigation within the document.
	Here are the summaries:`

	var composed_summaries string
	for _, fsch := range global_schemas {
		composed_summaries += fmt.Sprintf("File path: %v \n Summary: %v", fsch.path, fsch.summary)
	}

	return header + composed_summaries
}

func prompt_llm_model() {
	load_ollama_env()
	var full_endpoint string = strings.Join([]string{"http://", Ollama_endpoint, "/api/chat"}, "")

	prompt := compose_summaries_prompt()
	var msg message = message{
		Role:    "user",
		Content: prompt,
	}
	var req ollama_request = ollama_request{
		Model:    "llama3.1",
		Messages: []message{msg},
	}

	// Send the request
	result, err := send_request(full_endpoint, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save result to file
	err = save_result_to_file("README.md", result)
}

func save_result_to_file(filename, result string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(result)
	if err != nil {
		return err
	}

	fmt.Printf("Result saved to %s\n", filename)
	return nil
}
