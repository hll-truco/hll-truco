package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hll-truco/hll-truco/hll"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func MD5(data []byte) []byte {
	hasher := md5.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func Blake2b512(data []byte) []byte {
	hash := blake2b.Sum512(data)
	return hash[:]
}

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

func bytesToHexString(bytes []byte) string {
	hexString := hex.EncodeToString(bytes)
	return strings.TrimPrefix(hexString, "0x")
}

// SaveStringsToFile saves a slice of strings to a file, each string on a new line.
func SaveStringsToFile(filename string, lines []string) error {
	// Open the file for writing. Create it if it doesn't exist.
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	// Create a new writer.
	writer := bufio.NewWriter(file)

	// Write each string to the file, followed by a newline.
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("could not write to file: %w", err)
		}
	}

	// Flush the writer to ensure all data is written to the file.
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("could not flush writer: %w", err)
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

	// Blake2b512Dist := newDist(100)
	// for i := 0; i < n; i++ {
	// 	if i%reportEvery == 0 {
	// 		log.Println(i, float64(i)/float64(n))
	// 	}
	// 	hashSah256 := Blake2b512(toBytes(i))
	// 	_, val := hll.GetPosValDynamic(hashSah256, p)
	// 	Blake2b512Dist[int(val)-1] += 1
	// }
	// save(Blake2b512Dist, "Blake2b512dist-1B.json")

	// MD5Dist := newDist(100)
	// for i := 0; i < n; i++ {
	// 	if i%reportEvery == 0 {
	// 		log.Println(i, float64(i)/float64(n))
	// 	}
	// 	hashSah256 := MD5(toBytes(i))
	// 	_, val := hll.GetPosValDynamic(hashSah256, p)
	// 	MD5Dist[int(val)-1] += 1
	// }
	// save(MD5Dist, "md5dist-1B.json")

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
	// sha3_1024_Dist := newDist(100)
	// hashSah3_1024 := make([]byte, 128)
	// for i := 0; i < n; i++ {
	// 	if i%reportEvery == 0 {
	// 		log.Println(i, float64(i)/float64(n))
	// 	}
	// 	Sha3Optimized(hashSah3_1024, toBytes(i))
	// 	_, val := hll.GetPosValDynamic(hashSah3_1024, p)
	// 	sha3_1024_Dist[int(val)-1] += 1
	// }
	// save(sha3_1024_Dist, "sha3-1024-dist-1B.json")

	// hll.Sha3 1024
	newsha3_1024_Dist := newDist(100)
	fn := hll.NewSha3Hash(128)
	hashes := make([]string, 1_000_000)

	for i := 0; i < n; i++ {
		if i%reportEvery == 0 {
			log.Println(i, float64(i)/float64(n))
		}
		// compute hash
		fn.Write(toBytes(i))
		hashSah3_1024 := fn.Sum(nil)
		fn.Reset()

		if i < 1_000_000 {
			hashes[i] = bytesToHexString(hashSah3_1024)
		} else {
			break
		}

		_, val := hll.GetPosValDynamic(hashSah3_1024, p)
		newsha3_1024_Dist[int(val)-1] += 1
	}
	// save(newsha3_1024_Dist, "newsha3-1024-dist-1B.json")
	SaveStringsToFile("1M_go_sha3_1024_random.log", hashes)

	log.Println("total time", time.Since(start))
}

// bench

// md5 1B = 4m58.667003917s
// sha512 1B = 301 s = 5,017 mins
// blake512 1B = 6m50.655153375s
// sha3 1024 1B optimized = 14m43.369908542s
// newsha3 1B = 16m3.5948905s

// sha3 1024 4B (unoptimized) = 61 mins
