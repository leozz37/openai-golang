package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type request struct {
	Model       string  `json:"model"`
	Temperature float32 `json:"temperature"`
	Prompt      string  `json:"prompt"`
	Max_tokens  int32   `json:"max_tokens"`
}

type Response struct {
	Choices []struct {
		Text string `json:"text"`
	}
}

func Create(prompt, model string, max_tokens int32, temperature float32) (string, error) {
	instance_url := os.Getenv("OPENAI_INSTANCE_URL")
	bearer := "Bearer " + os.Getenv("OPENAI_API_KEY")

	body := request{
		Model:       model,
		Temperature: temperature,
		Prompt:      prompt,
		Max_tokens:  max_tokens,
	}

	json_data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", instance_url, bytes.NewBufferString(string(json_data)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	// Send request using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var inter Response

	err = json.Unmarshal(response, &inter)
	if err != nil {
		return "", err
	}

	return inter.Choices[0].Text, nil
}
