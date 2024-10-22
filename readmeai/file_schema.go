package main

import "fmt"

// As this is going to take all files from the repo and generate an AI summary from it
// I will need a data structure on how to store the semantic information that the AI Gave

type file_purpose int

const (
	documentation file_purpose = iota
	explanation
	implementation
	test
	other
)

type file_schema struct {
	path          string // Path to the file
	summary       string // Summary of the file
	semantic_info string // Semantic information of the file
	purpose       file_purpose
	useful_info   []string // Useful information that the AI found
}

func NewFileSchema(path string) {
	global_schemas = append(global_schemas, &file_schema{
		path: path,
	})
}

func (f *file_schema) Print() {
	fmt.Printf("Path: %v\nSummary: %v\nSemantic Info: %v\nPurpose: %v\nUseful Info: %v\n", f.path, f.summary, f.semantic_info, f.purpose, f.useful_info)
}

var global_schemas = []*file_schema{}
