package main

import (
"encoding/json"
"fmt"
"log"
"math/rand"
"net/http"
"os"
"strconv"
"time"
)

type InferenceRequest struct {
Input string `json:"input"`
}

type InferenceResponse struct {
Input       string  `json:"input"`
Prediction  string  `json:"prediction"`
Confidence  float64 `json:"confidence"`
Model       string  `json:"model"`
Instance    string  `json:"instance"`
ElapsedMs   int64   `json:"elapsedMs"`
}

func infer(input string) (string, float64) {
latencyMs := 150 + rand.Intn(150)
time.Sleep(time.Duration(latencyMs) * time.Millisecond)
classes := []string{"positive", "neutral", "negative"}
p := classes[rand.Intn(len(classes))]
conf := 0.6 + rand.Float64()*0.4
return p, conf
}

func handler(w http.ResponseWriter, r *http.Request) {
start := time.Now()
if r.Method != http.MethodPost {
http.Error(w, "POST only", http.StatusMethodNotAllowed)
return
}
var req InferenceRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, "bad json", http.StatusBadRequest)
return
}
pred, conf := infer(req.Input)
host, _ := os.Hostname()
resp := InferenceResponse{
Input:      req.Input,
Prediction: pred,
Confidence: conf,
Model:      "toy-sentiment-v1",
Instance:   host,
ElapsedMs:  time.Since(start).Milliseconds(),
}
w.Header().Set("Content-Type", "application/json")
_ = json.NewEncoder(w).Encode(resp)
}

func health(w http.ResponseWriter, r *http.Request) {
w.WriteHeader(http.StatusOK)
fmt.Fprintln(w, "ok")
}

func main() {
rand.Seed(time.Now().UnixNano())
port := 8080
if p := os.Getenv("PORT"); p != "" {
if n, err := strconv.Atoi(p); err == nil {
port = n
}
}
http.HandleFunc("/infer", handler)
http.HandleFunc("/healthz", health)
addr := fmt.Sprintf(":%d", port)
log.Printf("listening on %s\n", addr)
log.Fatal(http.ListenAndServe(addr, nil))
}
