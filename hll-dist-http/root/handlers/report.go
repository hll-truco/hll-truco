package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
)

// ReportRequest represents the request body for the /Report endpoint
type ReportRequest state.WorkerResult

// ReportHandler handles requests to the /Report endpoint
func ReportHandler(
	s *state.State,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding the request body into an ReportRequest object
		var reportReq ReportRequest
		if err := json.NewDecoder(r.Body).Decode(&reportReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddWorkerResult((*state.WorkerResult)(&reportReq))

		// Responding with status code 201
		w.WriteHeader(http.StatusCreated)
	}
}
