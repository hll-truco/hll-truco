package root

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/hll-truco/hll-truco/hll"
	"github.com/hll-truco/hll-truco/utils"
)

var (
	global  *hll.HyperLogLogExt = nil
	decoder *hll.HyperLogLogExt = nil
	total   uint64              = 0
	printer                     = utils.NewCronoPrinter(time.Second * 10)
)

func getNewExt() *hll.HyperLogLogExt {
	h1, err := hll.NewExt(16)
	if err != nil {
		panic(err)
	}
	return h1
}

func init() {
	// ini global
	global = getNewExt()
	// ini decode
	decoder = getNewExt()
}

// UpdateRequest represents the request body for the /update endpoint
type UpdateRequest struct {
	Gob string `json:"gob"`
}

func handleError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	slog.Error("DECODE_ERR")
}

// updateHandler handles requests to the /update endpoint
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	decoder.GobDecode(data)
	bump, err := global.Merge(decoder)

	if err != nil {
		handleError(err, w)
		return
	}

	total++

	if bump || printer.ShouldPrint() {
		if printer.ShouldPrint() {
			delta := printer.Check().Seconds()
			estimate := global.Count()
			slog.Info(
				"REPORT",
				"delta", delta,
				"estimate", estimate,
				"total", total)
		}
	}

	// Responding with status code 201
	w.WriteHeader(http.StatusCreated)
}
