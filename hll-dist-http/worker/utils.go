package worker

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendPOSTJsonData(url string, data any) {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("JSON_ERR", "error", err)
		return
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("HTTP_CREATE_ERR", "error", err)
		return
	}

	// Set the appropriate content type
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("HTTP_SEND_ERR", "error", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 201 Created
	if resp.StatusCode != http.StatusCreated {
		slog.Error("UNEXPECTED_RES", "status", resp.StatusCode)
	}
}
