package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type InferenceStatus struct {
	Status          string      `json:"status"`
	RuntimeMs       int         `json:"runtime_ms"`
	Cost            float64     `json:"cost"`
	TokensGenerated interface{} `json:"tokens_generated"`
	TokensInput     interface{} `json:"tokens_input"`
}

type Response struct {
	RequestID           string          `json:"request_id"`
	InferenceStatus     InferenceStatus `json:"inference_status"`
	Images              []string        `json:"images"`
	NSFWContentDetected []bool          `json:"nsfw_content_detected"`
	Seed                int64           `json:"seed"`
}

func main() {
	url := "https://api.deepinfra.com/v1/inference/black-forest-labs/FLUX-1-dev"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("prompt", "A front-facing view of a pixelated F1 car, with exaggerated features and glowing headlights, surrounded by a simple shadowy background.")
	err := writer.Close()
	if err != nil {
		fmt.Println("Error creating payload:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Authorization", "bearer 4uqGj2dDuN5LKJdFv6Cd5A5aSVdradD2")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	if len(response.Images) == 0 {
		fmt.Println("No images found in response")
		return
	}

	image := response.Images[0]

	decodeBase64Image(image)
	fmt.Println(response.NSFWContentDetected)
}

func decodeBase64Image(base64String string) {
	base64String = strings.TrimPrefix(base64String, "data:image/png;base64,")

	// Decode the Base64 string
	decodedData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err)
		return
	}

	// Write the decoded data to a file
	err = ioutil.WriteFile("output.png", decodedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Image successfully decoded and saved as output.png")
}
