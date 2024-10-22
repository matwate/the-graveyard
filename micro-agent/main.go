package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prompt := `How can we represent the rule "There has to be 8 queens on the board" using first order logic, and a descriptor that describes that a queen is in the position x,y such as R(X,Y) <=> True if a queen is on the position X,Y `

	ending := "Generate a final conclusion of the solution, extracted away from any expert commentaries, and the full logical statement"

	ToT_starter := make_ToT_starter(prompt)
	messages := []string{ToT_starter}

	// Define the number of iterations
	numIterations := 5

	for i := 0; i < numIterations; i++ {
		chat_messages := chatWithAI(messages)
		req := make_request(chat_messages)
		ollama_response := send_request_to_ollama(req)
		messages = append(messages, ollama_response.Message.Content, "Expand on your previous response by considering the key aspects of the topic and provide a new perspective, taking into account any relevant context or information provided.")
		fmt.Printf("Iteration %d completed\n", i+1)
	}

	messages = messages[:len(messages)-1]

	messages = append(messages, ending)
	chat_messages := chatWithAI(messages)
	req := make_request(chat_messages)
	res := send_request_to_ollama(req)
	fmt.Println(res.Message.Content)
	messages = append(messages, res.Message.Content)

	// Save the chat to a file
	saveChatToFile(messages, "chat_log.txt")
	saveResultToFile(messages[len(messages)-1], "result_log.txt")
}

func saveChatToFile(messages []string, filename string) {
	// Join the messages into a single string
	chatContent := strings.Join(messages, "\n\n")

	// Write the string to a file
	err := os.WriteFile(filename, []byte(chatContent), 0644)
	if err != nil {
		fmt.Printf("Error saving chat to file: %v\n", err)
	} else {
		fmt.Printf("Chat saved to file: %s\n", filename)
	}
}

func saveResultToFile(message, filename string) {
	err := os.WriteFile(filename, []byte(message), 0644)
	if err != nil {
		fmt.Printf("Error saving result to file: %v\n", err)
	} else {
		fmt.Printf("Result saved to file: %s\n", filename)
	}
}
