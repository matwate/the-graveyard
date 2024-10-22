package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var Ollama_endpoint string

type ollama_request struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Message            Message   `json:"message"`
	Done               bool      `json:"done"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int64     `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func load_ollama_env() (err error) {
	err = godotenv.Load()
	if err != nil {
		return err
	}
	Ollama_endpoint = os.Getenv("OLLAMA_ENDPOINT")
	return nil

}

func fill_schema_with_llm(fsch *file_schema) (err error) {
	load_ollama_env()
	var full_endpoint string = strings.Join([]string{"http://", Ollama_endpoint, "/api/chat"}, "")
	content, err := prompt_for_file_schema(fsch)
	if err != nil {
		return err
	}
	var msg message = message{
		Role:    "user",
		Content: content,
	}
	var req ollama_request = ollama_request{
		Model:    "llama3.1",
		Messages: []message{msg},
	}
	result, err := send_request(full_endpoint, req)
	if err != nil {
		return err
	}
	fsch.summary = result
	return nil

}

func prompt_for_file_schema(fsch *file_schema) (content string, err error) {
	file, err := os.Open(fsch.path)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v with error: %v", fsch.path, err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	var code_string string
	for reader.Scan() {
		code_string += reader.Text()
	}
	content = "Right now, you are in the process of generating info per file  for a github repo, i will provide you this file, that is only one of the many composing the entire repo, you must generate the following things for this code file: Summary, Semantic info, Purpose(Out of this options: documentation, explanation, implementation, test, other), and a bullet point list of useful information of this file. Each section should be clearly separated by two newlines so i can parse it programatically. Here's the contents:\n" + code_string
	return content, nil
}

func send_request(endpoint string, req ollama_request) (result string, err error) {
	// Marshal the request struct to JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
	}
	fmt.Println(string(jsonData))

	// Create a new HTTP POST request
	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to send request, status code: %d", resp.StatusCode)
	}

	// Read the response body;
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response body into the Response struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Message.Content, nil

}
