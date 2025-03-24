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
    fmt.Println("🚀 Listening on :8443")
    log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
}