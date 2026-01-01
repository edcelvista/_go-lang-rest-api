package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // automatically expose /debug/pprof/
	"os"

	router "pkg/router"
	// CPU Profiling: /debug/pprof/profile?seconds=30 (takes a 30-second snapshot)
	// Goroutines: /debug/pprof/goroutine
	// Heap (memory): /debug/pprof/heap
	// Block: /debug/pprof/block
	// OR
	// go tool pprof -http=localhost:6061 "http://localhost:6060/debug/pprof/profile?seconds=30"
	// go tool pprof http://localhost:6060/debug/pprof/heap
	// # wait 1–5 minutes
	// go tool pprof http://localhost:6060/debug/pprof/heap
	// go tool pprof -diff_base=heap1.pprof heap2.pprof
	// go tool pprof -http=localhost:6061 "http://localhost:6060/debug/pprof/heap?gc=1" # force GC to confirm leak, if memory stays means leaks
	// go tool pprof http://localhost:6060/debug/pprof/goroutine # check goroutine leaks common cause
	// Requires brew install graphviz
)

func main() {
	if os.Getenv("APP_ENV") == "development" {
		log.Println("👮‍♂️ Enabling pprof for profiling")
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	router.Run()
}
