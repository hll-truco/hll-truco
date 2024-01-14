package utils

import (
	"log"
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
