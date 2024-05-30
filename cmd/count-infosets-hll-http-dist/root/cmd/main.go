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

// global state
var (
	exitChan = make(chan bool)
)

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
	mux := http.NewServeMux()

	// Create the server with the mux
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *portFlag),
		Handler: mux,
	}

	mux.HandleFunc("/version", root.VersionHandler)
	mux.HandleFunc("/update", root.UpdateHandler)
	mux.HandleFunc("/exit", root.ExitHandler(server, exitChan))

	go func() {
		slog.Info("UP", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("BIND_ERR", "addr", server.Addr, "err", err)
			panic(err)
		}
	}()

	<-exitChan // Wait for signal to exit
	slog.Info("DOWN")
	os.Exit(0) // Exit the program
}
