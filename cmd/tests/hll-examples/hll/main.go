package main

import (
	"flag"
	"fmt"
	"hash"
	"log/slog"
	"os"
	"time"

	"github.com/hll-truco/hll-truco/hll"
	"github.com/hll-truco/hll-truco/hll-dist-http/worker"
	"github.com/hll-truco/hll-truco/utils"
)

// flags/parametros:
var (
	nFlag         = flag.Uint64("n", 100_000, "Cardinality to estimate")
	limitFlag     = flag.Int("limit", 60, "Run time limit (in seconds) (default 1m)")
	hashIDFlag    = flag.String("hash", "sha160", "Infoset hashing function")
	reportFlag    = flag.Int("report", 10, "Delta (in seconds) for printing log msgs")
	precisionFlag = flag.Int("precision", 16, "HLL precision (defaults to 16)")
)

func init() {
	flag.Parse()

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// dump args
	slog.Info(
		"START",
		"n", *nFlag,
		"limit", *limitFlag,
		"hash", *hashIDFlag,
		"report", *reportFlag,
		"precision", *precisionFlag)
}

// var byteSlice = make([]byte, 8)

// 1_000_000_000
//   258_175_523

func Hashify(h hash.Hash, x uint64) []byte {
	// bytes conversion
	// binary.LittleEndian.PutUint64(byteSlice, x)
	byteSlice := []byte(fmt.Sprintf("%d", x))
	// calculate hash
	h.Reset()
	h.Write(byteSlice)
	return h.Sum(nil)
}

func main() {
	// max ui64 ~ 1.8e19

	// limit := time.Second * time.Duration(*limitFlag)
	precision := uint8(*precisionFlag)
	h, _ := hll.NewExt(precision)
	hashFn := worker.ParseHashFn(*hashIDFlag)
	printer := utils.NewCronoPrinter(time.Second * time.Duration(*reportFlag))
	start := time.Now()
	nodesVisited := uint64(0)

	// alternative time-limited stop condition
	// time.Since(start) < limit

	for nodesVisited = 0; nodesVisited < *nFlag; nodesVisited++ {
		_hash := Hashify(hashFn, nodesVisited)
		h.Add(_hash)
		if printer.ShouldPrint() {
			delta := printer.Check().Seconds()
			// e := h.CountBig()
			e := h.Count()
			p := float64(nodesVisited) / float64(*nFlag)
			slog.Info(
				"REPORT",
				"delta", delta,
				"visited", nodesVisited,
				"estimate", e,
				"progress", fmt.Sprintf("%.2f", p))
		}
	}

	slog.Info(
		"RESULTS",
		"finalEstimate", h.CountBig(),
		"nodesVisited:", nodesVisited,
		"finished", time.Since(start).Seconds())
}
