package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/hll-truco/hll-truco/hll-dist-http/root/handlers"
	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
	"github.com/hll-truco/hll-truco/utils"
)

// flags/parametros:
var (
	portFlag   = flag.Int("port", 8080, "HTTP port")
	reportFlag = flag.Int("report", 1, "Delta (in seconds) for printing log msgs")
	saveFlag   = flag.String("save", "", "Full path to file to save all progress on exit")
	resumeFlag = flag.String("resume", "", "Full path to file to save all progress on exit")
)

func init() {
	flag.Parse()

	// always save
	if noSave := len(*saveFlag) == 0; noSave {
		*saveFlag = fmt.Sprintf("%d.json", rand.Int63())
	}

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info(
		"START",
		"port", *portFlag,
		"resume", *resumeFlag,
		"save", *saveFlag)
}

func main() {
	var (
		state    = state.NewState()
		printer  = utils.NewCronoPrinter(time.Second * time.Duration(*reportFlag))
		mux      = http.NewServeMux()
		exitChan = make(chan bool)
		server   = &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%d", *portFlag),
			Handler: mux,
		}
	)

	if len(*resumeFlag) > 0 {
		state.Load(*resumeFlag)
		slog.Info(
			"LOADED",
			"file", *resumeFlag,
			"total", state.Total,
			"estimate", state.Estimate())
	}

	mux.HandleFunc("/version", handlers.VersionHandler)
	mux.HandleFunc("/update", handlers.UpdateHandler(state, printer))
	mux.HandleFunc("/report", handlers.ReportHandler(state))
	mux.HandleFunc("/exit", handlers.ExitHandler(server, exitChan, state, printer, *saveFlag))

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

go run cmd/count-infosets-hll-dist-http/root/main.go -port=8080 | tee logs/hll-dist-http/example.log

*/
