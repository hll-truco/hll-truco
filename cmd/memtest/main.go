package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/hll-truco/experiments/utils"
)

var printer *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 1)

/*

- HeapAlloc: This is the number of bytes allocated on the heap and still in use
  by the Go application. This value increases as you allocate memory (e.g.,
  create new objects) and decreases as the garbage collector frees up memory
  that’s no longer in use.

- TotalAlloc: This is the cumulative count of bytes allocated on the heap over
  the lifetime of the program. Unlike HeapAlloc, this value only increases - it
  does not decrease when memory is freed. This means that TotalAlloc gives you
  the total amount of memory allocated by the program, regardless of whether it
  has been freed since.

- Sys: This is the total amount of memory obtained from the operating system by
  the Go runtime. It includes memory used by the Go runtime itself as well as
  memory allocated on the heap (regardless of whether it’s currently in use or
  has been returned to the OS). It’s worth noting that this value will be larger
  than the total heap allocations because it includes overheads and other
  structures managed by the Go runtime

*HeapAlloc*

run `python cmd/memtest/plot_dynm_vs_fixed.py` and judge by yourself

From https://go.dev/blog/go119runtime

"
(...)
The first is that when the peak memory use of an application is unpredictable,
GOGC alone offers virtually no protection from running out of memory. With just
GOGC, the Go runtime is simply unaware of how much memory it has available to
it. Setting a memory limit enables the runtime to be robust against transient,
recoverable load spikes by making it aware of when it needs to work harder to
reduce memory overhead.
(...)
"


*/

func getMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("HeapAlloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v",
		bToMb(m.HeapAlloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

var delay = time.Millisecond * 5

func DynamicSlice(n int) []byte {
	data := make([]byte, 0)
	for i := 0; i < n; i++ {
		data = append(data, make([]byte, 1024*1024)...)
		printer.Print(getMemUsage())
		time.Sleep(delay)
	}
	return data
}

func FixedSlice(n int) []byte {
	data := make([]byte, 0, n*1024*1024)
	for i := 0; i < n; i++ {
		data = append(data, make([]byte, 1024*1024)...)
		printer.Print(getMemUsage())
		time.Sleep(delay)
	}
	return data
}

func main() {
	fmt.Println("GOMEMLIMIT:", os.Getenv("GOMEMLIMIT"))
	fmt.Printf("Current process ID: %d\n", os.Getpid())

	var n int
	flag.IntVar(&n, "n", 0, "Amount of memory to fill in MiB")
	flag.Parse()

	data := DynamicSlice(n)
	// data := FixedSlice(n)

	fmt.Println("done. sleeping 10s.", len(data))
	time.Sleep(1 * time.Minute)
	fmt.Println(len(data))
}
