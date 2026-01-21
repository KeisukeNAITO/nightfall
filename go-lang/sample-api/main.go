package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type helloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := helloResponse{Message: "hello world"}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

type doubleResponse struct {
	Input  int64 `json:"input"`
	Double int64 `json:"double"`
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	const prefix = "/api/v1/double/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.NotFound(w, r)
		return
	}

	numberStr := strings.TrimPrefix(r.URL.Path, prefix)
	if numberStr == "" {
		http.Error(w, "number is required", http.StatusBadRequest)
		return
	}

	n, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	// overflow check: n*2 within int64
	if n > math.MaxInt64/2 || n < math.MinInt64/2 {
		http.Error(w, "number out of range", http.StatusBadRequest)
		return
	}

	resp := doubleResponse{Input: n, Double: n * 2}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/hello", helloHandler)
	mux.HandleFunc("/api/v1/double/", doubleHandler)

	addr := ":8080"
	log.Printf("starting hello-api server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
