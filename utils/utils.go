package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/adler32"
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
	start     time.Time
	lastPrint time.Time
	delta     time.Duration
}

func (p *CronoPrinter) Print(x string) {
	if time.Since(p.lastPrint) > p.delta {
		p.lastPrint = time.Now()
		log.Printf("(%v) %s\n", time.Since(p.start), x)
	}
}

func NewCronoPrinter(delta time.Duration) *CronoPrinter {
	return &CronoPrinter{
		start:     time.Now(),
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
