package main

import (
	_ "net/http/pprof"
	"net/http"
	"time"
)

func heavyHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	b := make([]byte, 100*1024*1024) // 100MB
	for i := range b {
		b[i] = byte(i)
	}
	w.Write([]byte("done"))
}

func main() {
	http.HandleFunc("/heavy", heavyHandler)
	go http.ListenAndServe("localhost:6060", nil) // pprof
	http.ListenAndServe(":8080", nil)
}
