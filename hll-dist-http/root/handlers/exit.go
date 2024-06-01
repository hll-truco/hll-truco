package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
	"github.com/hll-truco/hll-truco/utils"
)

// Handler for /exit endpoint
func ExitHandler(
	server *http.Server,
	exitChan chan bool,
	state *state.State,
	crono *utils.CronoPrinter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shutting down the server..."))
		go func() {
			time.Sleep(1 * time.Second) // Give the response some time to be sent
			if err := server.Shutdown(context.Background()); err != nil {
				slog.Error(
					"SHUTDOWN_ERR",
					"err", err)
			}
			state.Results()
			exitChan <- true // Signal the main function to exit
		}()
	}
}
