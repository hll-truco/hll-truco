package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/adler32"
	"log"
	"os"
	"runtime"
	"strconv"
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
	start     time.Time
	lastPrint time.Time
	delta     time.Duration
}

func (p *CronoPrinter) ShouldPrint() bool {
	return time.Since(p.lastPrint) > p.delta
}

func (p *CronoPrinter) Check() time.Duration {
	p.lastPrint = time.Now()
	return time.Since(p.start)
}

func NewCronoPrinter(delta time.Duration) *CronoPrinter {
	return &CronoPrinter{
		start:     time.Now(),
		lastPrint: time.Now(),
		delta:     delta,
	}
}

type MemUsage struct {
	HeapAlloc  uint64
	TotalAlloc uint64
	Sys        uint64
	NumGC      uint32
}

func (mu *MemUsage) String() string {
	return fmt.Sprintf(
		"HeapAlloc=%v MiB\tTotalAlloc=%v MiB\tSys = %v MiB\tNumGC=%v",
		ByteToMb(mu.HeapAlloc),
		ByteToMb(mu.TotalAlloc),
		ByteToMb(mu.Sys),
		mu.NumGC)
}

// returns memory usage in MiB
func GetMemUsage() *MemUsage {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &MemUsage{
		HeapAlloc:  ByteToMb(m.HeapAlloc),
		TotalAlloc: ByteToMb(m.TotalAlloc),
		Sys:        ByteToMb(m.Sys),
		NumGC:      m.NumGC,
	}
}

func ByteToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func ParseHashFn(ID string) hash.Hash {
	switch ID {
	case "sha160":
		return sha1.New()
	case "sha256":
		return sha256.New()
	case "sha512":
		return sha512.New()
	case "adler32":
		return adler32.New()
	}
	log.Panicf("hash `%s` not found", ID)
	return nil
}

func TimeLimitReached(started time.Time, envvar string) bool {
	limitStr := os.Getenv(envvar)
	if limitStr == "" || limitStr == "-1" {
		return false
	}

	limitSeconds, err := strconv.Atoi(limitStr)
	if err != nil {
		return false
	}

	elapsed := time.Since(started).Seconds()
	return elapsed > float64(limitSeconds)
}
