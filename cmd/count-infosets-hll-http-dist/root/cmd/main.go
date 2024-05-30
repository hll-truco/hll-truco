package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/hll-truco/hll-truco/cmd/count-infosets-hll-http-dist/root"
)

// flags/parametros:
var (
	portFlag = flag.Int("port", 8080, "HTTP port")
)

// global vars
// var (
// 	deck []int = nil
// )

func init() {
	flag.Parse()

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info(
		"START",
		"port", *portFlag)
}

func main() {
	// Registering handlers for /version and /update endpoints
	http.HandleFunc("/version", root.VersionHandler)
	http.HandleFunc("/update", root.UpdateHandler)

	// Starting the server
	addr := fmt.Sprintf("0.0.0.0:%d", *portFlag)
	slog.Info("START", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("SERVER_ERR", "msg", err)
	}
}
