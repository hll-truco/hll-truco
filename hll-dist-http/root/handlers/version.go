package handlers

import (
	"encoding/json"
	"net/http"
)

const version = "0.0.0"

// VersionResponse represents the response for the /version endpoint
type VersionResponse struct {
	Version string `json:"version"`
}

// versionHandler handles requests to the /version endpoint
func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	// Creating a VersionResponse object
	versionRes := VersionResponse{
		Version: version,
	}

	// Encoding the VersionResponse object to JSON
	jsonResponse, err := json.Marshal(versionRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")

	// Writing the JSON response with status code 200
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
