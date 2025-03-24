package main

import (
    "fmt"
    "log"
    "net/http"

    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("📥 Got request from:", r.RemoteAddr, "proto:", r.Proto)
    fmt.Fprintf(w, "Hello from Go!\n")
}

func main() {
    h := http.HandlerFunc(handler)

    server := &http.Server{
        Addr: ":8080",
        Handler: h2c.NewHandler(h, &http2.Server{}),
    }

    fmt.Println("🚀 Listening on :8080 with h2c support")
    log.Fatal(server.ListenAndServe())
}
