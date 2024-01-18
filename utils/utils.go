package utils

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// 34_320
// que imprima cada 1_000
func PrintEvery(x uint64, delta uint64) {
	if x%delta == 0 {
		log.Println(x)
	}
}

type CronoPrinter struct {
	lastPrint time.Time
	delta     time.Duration
}

func (p *CronoPrinter) Print(x string) {
	if time.Since(p.lastPrint) > p.delta {
		p.lastPrint = time.Now()
		log.Println(x)
	}
}

func NewCronoPrinter(delta time.Duration) *CronoPrinter {
	return &CronoPrinter{
		lastPrint: time.Now(),
		delta:     delta,
	}
}

func GetMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("HeapAlloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v",
		ByteToMb(m.HeapAlloc),
		ByteToMb(m.TotalAlloc),
		ByteToMb(m.Sys),
		m.NumGC)
}

func ByteToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
