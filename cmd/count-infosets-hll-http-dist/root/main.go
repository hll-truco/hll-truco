package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/hll-truco/hll-truco/hll-dist-http/root/handlers"
	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
	"github.com/hll-truco/hll-truco/utils"
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

func main() {
	var (
		state    = state.NewState()
		printer  = utils.NewCronoPrinter(time.Second * 10)
		mux      = http.NewServeMux()
		exitChan = make(chan bool)
		server   = &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%d", *portFlag),
			Handler: mux,
		}
	)

	mux.HandleFunc("/version", handlers.VersionHandler)
	mux.HandleFunc("/update", handlers.UpdateHandler(state, printer))
	mux.HandleFunc("/report", handlers.ReportHandler(state))
	mux.HandleFunc("/exit", handlers.ExitHandler(server, exitChan, state, printer))

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

/*

E.g.,

go run cmd/count-infosets-hll-http-dist/root/main.go -port=8080 | tee logs/hll-dist-http/example.log

*/
