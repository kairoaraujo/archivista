package tuf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AddArtifact sends an HTTP POST request to add an artifact
func AddArtifact(url string, payload ArtifactPayload) error {
	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// Add any additional headers or authentication if needed

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	// Handle the response as needed

	return nil
}
