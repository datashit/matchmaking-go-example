package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/datashit/matchmaking-go"
)

func wsHand(w http.ResponseWriter, r *http.Request) {
	matchmaking.Accept(w, r)
}

func main() {
	fmt.Printf("CPU - %v", runtime.GOMAXPROCS(runtime.NumCPU())) //Using All Core
	matchmaking.CreateProcesWorker(10)
	http.HandleFunc("/", wsHand)
	http.ListenAndServe(":80", nil)
}
