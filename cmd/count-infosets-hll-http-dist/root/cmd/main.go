package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/hll-truco/hll-truco/cmd/count-infosets-hll-http-dist/root"
)

// flags/parametros:
var (
	portFlag = flag.Int("port", 8080, "HTTP port")
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

// Handler for /exit endpoint
func ExitHandler(server *http.Server, exitChan chan bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shutting down the server..."))
		go func() {
			time.Sleep(1 * time.Second) // Give the response some time to be sent
			if err := server.Shutdown(context.Background()); err != nil {
				fmt.Println("Error shutting down the server:", err)
			}
			exitChan <- true // Signal the main function to exit
		}()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", root.VersionHandler)
	mux.HandleFunc("/update", root.UpdateHandler)

	// Create the server with the mux
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *portFlag),
		Handler: mux,
	}

	exitChan := make(chan bool)

	// Add the /exit handler
	mux.HandleFunc("/exit", ExitHandler(server, exitChan))

	go func() {
		slog.Info("SERVER_UP", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	<-exitChan // Wait for signal to exit
	slog.Info("SERVER_DOWN")
	os.Exit(0) // Exit the program
}
