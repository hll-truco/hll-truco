package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
	"github.com/hll-truco/hll-truco/utils"
)

// UpdateRequest represents the request body for the /update endpoint
type UpdateRequest struct {
	Gob string `json:"gob"`
}

// updateHandler handles requests to the /update endpoint
func UpdateHandler(
	s *state.State,
	crono *utils.CronoPrinter,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding the request body into an UpdateRequest object
		var updateReq UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := base64.StdEncoding.DecodeString(updateReq.Gob)
		if err != nil {
			handleError(err, w)
			return
		}

		bump, err := s.Merge(data)

		if err != nil {
			handleError(err, w)
			return
		}

		if bump || crono.ShouldPrint() {
			if crono.ShouldPrint() {
				_ = crono.Check().Seconds()
				s.Report()
			}
		}

		// Responding with status code 201
		w.WriteHeader(http.StatusCreated)
	}
}
