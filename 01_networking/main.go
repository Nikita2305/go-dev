package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("📥 Got request from:", r.RemoteAddr)
    fmt.Fprintf(w, "Hello from Go!\n")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("🚀 Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}