package worker

import (
	"bytes"
	"encoding/json"
	"hash"
	"log/slog"
	"math/rand"
	"net/http"

	"github.com/hll-truco/hll-truco/hll"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/gotruco/pdt"
)

func UniformPick(chis [][]pdt.IJugada) pdt.IJugada {
	// hago un flatten del vector chis
	n := len(chis) * 15
	flatten := make([]pdt.IJugada, 0, n)

	for _, chi := range chis {
		flatten = append(flatten, chi...)
	}

	// elijo una jugada al azar
	rfix := rand.Intn(len(flatten))

	return flatten[rfix]
}

func ParseHashFn(hashFn string) hash.Hash {
	if hashFn == "sha3" {
		slog.Warn("USING_SHA3SHAKE", "size", 128)
		return hll.NewSha3Hash(128)
	} else {
		slog.Warn("USING_FIXED_HASH", "hash", hashFn)
		return utils.ParseHashFn(hashFn)
	}
}

func sendPOSTJsonData(url string, data any) {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("JSON_ERR", "error", err)
		return
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("HTTP_CREATE_ERR", "error", err)
		return
	}

	// Set the appropriate content type
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("HTTP_SEND_ERR", "error", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 201 Created
	if resp.StatusCode != http.StatusCreated {
		slog.Error("UNEXPECTED_RES", "status", resp.StatusCode)
	}
}
