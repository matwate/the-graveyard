package main

import "fmt"

func make_ToT_starter(prompt string) string {
	return fmt.Sprintf(
		`Imagine three different experts are solving a problem
		All experts will write down 1 step of their thinking,
		then share it with the group.
		Then all experts will go on to the next step, etc.
		If any expert realises they're wrong at any point then they leave.
		The problem is.. %v`,
		prompt,
	)
}

func make_request(messages []message) ollama_request {
	return ollama_request{
		Model:    "llama3.1",
		Messages: messages,
		Stream:   false,
	}
}

func chatWithAI(messages []string) []message {
	// Implementation for a sequence of messages back and forth between the user and the AI
	var output []message
	for i, v := range messages {
		switch i % 2 {
		case 1:
			output = append(output,
				message{
					Role:    "assistant",
					Content: v,
				},
			)
		case 0:
			output = append(output,
				message{
					Role:    "user",
					Content: v,
				},
			)
		}
	}
	return output

}
