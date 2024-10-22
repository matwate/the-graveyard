package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Notebook struct {
	Cells         []Cell   `json:"cells"`
	Metadata      Metadata `json:"metadata"`
	NBFormat      int      `json:"nbformat"`
	NBFormatMinor int      `json:"nbformat_minor"`
}

type Cell struct {
	CellType string   `json:"cell_type"`
	Source   []string `json:"source"`
}

type Metadata struct {
	KernelSpec KernelSpec `json:"kernelspec"`
}

type KernelSpec struct {
	DisplayName string `json:"display_name"`
	Language    string `json:"language"`
	Name        string `json:"name"`
}

func main() {
	// Read and parse the notebook file
	notebookData, err := os.ReadFile("notebook.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var notebook Notebook
	err = json.Unmarshal(notebookData, &notebook)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Process and execute all code cells
	executeAllCells(notebook.Cells)
}

func executeAllCells(cells []Cell) {
	// Prepare the complete C++ program
	programCode := "#include <iostream>\n#include <vector>\n#include <string>\n#include <algorithm>\n\n"
	programCode += "int main() {\n"

	for _, cell := range cells {
		switch cell.CellType {
		case "markdown":
			printMarkdownCell(cell)
		case "code":
			programCode += "\n// --- Begin Cell ---\n"
			programCode += strings.Join(cell.Source, "\n")
			programCode += "\n// --- End Cell ---\n"
		}
	}

	programCode += "\nreturn 0;\n}"

	tmpfile, err := os.CreateTemp("", "notebook_*.cpp")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(programCode)); err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}
	if err := tmpfile.Close(); err != nil {
		fmt.Println("Error closing temporary file:", err)
		return
	}

	// Compile the program
	compiledFile := strings.TrimSuffix(tmpfile.Name(), ".cpp")
	cmd := exec.Command("g++", "", tmpfile.Name(), "-o", compiledFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Compilation error:", string(output))
		return
	}
	defer os.Remove(compiledFile)

	// Run the compiled program
	cmd = exec.Command(compiledFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Runtime error:", err)
	}
	fmt.Println("Program output:")
	fmt.Println(string(output))
}

func printMarkdownCell(cell Cell) {
	fmt.Println("Markdown Cell:")
	for _, line := range cell.Source {
		fmt.Println(line)
	}
	fmt.Println()
}
