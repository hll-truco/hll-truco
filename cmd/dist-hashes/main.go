package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hll-truco/hll-truco"
	"golang.org/x/crypto/sha3"
)

func Sha3(data []byte) []byte {
	hash1024bits := make([]byte, 128)
	sha3.ShakeSum256(hash1024bits, data)
	return hash1024bits
}

func Sha3Optimized(dest []byte, data []byte) {
	sha3.ShakeSum256(dest, data)
}

func Sha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func Sha512(data []byte) []byte {
	hash := sha512.Sum512(data)
	return hash[:]
}

func toBytes(x int) []byte {
	return []byte(fmt.Sprintf("%d", x))
}

// newDist returns a map (dictionary) where its keys span from 0 to n,
// and all its values are initialized to 0
func newDist(n int) map[int]int {
	m := make(map[int]int)
	for i := 0; i <= n; i++ {
		m[i] = 0
	}
	return m
}

func save(data any, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: Set indentation for readability

	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	p := uint8(16)
	n := 1_000_000_000
	reportEvery := int(float64(n) * 0.01) // every 1% percent
	start := time.Now()

	// sha256
	// sha256Dist := newDist(100)
	// for i := 0; i < n; i++ {
	// 	hashSah256 := Sha256(toBytes(i))
	// 	_, val := hll.GetPosValDynamic(hashSah256, p)
	// 	// t.Logf("i:%d zeros+1:%d sha256: %v\n", i, val, hashSah256)
	// 	sha256Dist[int(val)-1] += 1
	// }
	// save(sha256Dist, "sha256dist-1B.json")

	// sha512
	// sha512Dist := newDist(100)
	// for i := 0; i < n; i++ {
	// 	if i%reportEvery == 0 {
	// 		log.Println(i, float64(i)/float64(n))
	// 	}
	// 	hashSah256 := Sha512(toBytes(i))
	// 	_, val := hll.GetPosValDynamic(hashSah256, p)
	// 	sha512Dist[int(val)-1] += 1
	// }
	// save(sha512Dist, "sha512dist-1B.json")

	// sha3 1024
	// sha3_1024_Dist := newDist(100)
	// for i := 0; i < n; i++ {
	// 	if i%reportEvery == 0 {
	// 		log.Println(i, float64(i)/float64(n))
	// 	}
	// 	hashSah3_1024 := Sha3(toBytes(i))
	// 	_, val := hll.GetPosValDynamic(hashSah3_1024, p)
	// 	sha3_1024_Dist[int(val)-1] += 1
	// }
	// save(sha3_1024_Dist, "sha3-1024-dist-1B.json")

	// Sha3Optimized 1024
	sha3_1024_Dist := newDist(100)
	hashSah3_1024 := make([]byte, 128)
	for i := 0; i < n; i++ {
		if i%reportEvery == 0 {
			log.Println(i, float64(i)/float64(n))
		}
		Sha3Optimized(hashSah3_1024, toBytes(i))
		_, val := hll.GetPosValDynamic(hashSah3_1024, p)
		sha3_1024_Dist[int(val)-1] += 1
	}
	save(sha3_1024_Dist, "sha3-1024-dist-1B.json")

	log.Println("total time", time.Since(start))
}

// bench
// sha3 1024 4B (unoptimized) = 61 mins
// sha512 1B = 301 s = 5,017 mins
// sha3 1024 1B optimized = 14m43.369908542s
