package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
)

// ReportRequest represents the request body for the /Report endpoint
type ReportRequest struct {
	Gob string `json:"gob"`
}

// ReportHandler handles requests to the /Report endpoint
func ReportHandler(
	state *state.State,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding the request body into an ReportRequest object
		var ReportReq ReportRequest
		if err := json.NewDecoder(r.Body).Decode(&ReportReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := base64.StdEncoding.DecodeString(ReportReq.Gob)
		if err != nil {
			handleError(err, w)
			return
		}

		state.Decoder.GobDecode(data)
		_, err = state.Global.Merge(state.Decoder)

		if err != nil {
			handleError(err, w)
			return
		}

		state.Total++

		// Responding with status code 201
		w.WriteHeader(http.StatusCreated)
	}
}
