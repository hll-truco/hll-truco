package root

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

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
