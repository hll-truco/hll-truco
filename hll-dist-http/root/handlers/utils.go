package handlers

import (
	"log/slog"
	"net/http"
)

func handleError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	slog.Error("DECODE_ERR")
}
