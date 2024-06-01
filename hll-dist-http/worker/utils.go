package worker

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
)

type UpdateRequest struct {
	Gob string `json:"gob"`
}

func SendUpdateRequest(baseURL string, gobString string) {
	url := baseURL + "/update"
	// Create the UpdateRequest struct
	update := UpdateRequest{Gob: gobString}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(update)
	if err != nil {
		slog.Warn("JSON_ERR", "error", err)
		return
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Warn("HTTP_CREATE_ERR", "error", err)
		return
	}

	// Set the appropriate content type
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Warn("HTTP_SEND_ERR", "error", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 201 Created
	if resp.StatusCode != http.StatusCreated {
		slog.Warn("UNEXPECTED_RES", "status", resp.StatusCode)
	}
}
