package main

import (
	"flag"
	"fmt"
	"os"

	"grpc-go-scratch/goroutine"
)

func main() {
	var (
		demo     string
		profiler bool
	)

	flag.StringVar(&demo, "demo", "", "Select demo to run: buffer-channel, worker-pool, unbounded")
	flag.BoolVar(&profiler, "profiler", false, "Enable pprof profiler on localhost:6060")
	flag.Parse()

	if profiler {
		goroutine.InitProfiler()
	}

	switch demo {
	case "buffer-channel":
		fmt.Println("Running Buffer Channel Demo...")
		goroutine.RunBufferChannelDemo()
	case "worker-pool":
		fmt.Println("Running Worker Pool Demo...")
		goroutine.RunWorkerPoolDemo()
	case "unbounded":
		fmt.Println("Running Unbounded Goroutines Demo (not recommended)...")
		goroutine.RunUnboundedGoroutinesDemo()
	case "memory-reuse":
		fmt.Println("Running Memory Reuse Demo...")
		goroutine.MemoryReuse()
	case "stream-chunk":
		fmt.Println("Running Stream Chunk Demo...")
		goroutine.StreamChunks()
	default:
		fmt.Println("Usage: go run main.go -demo <demo-name>")
		fmt.Println("Available demos:")
		fmt.Println("  buffer-channel : Demonstrates buffered channel pattern for producer-consumer")
		fmt.Println("  worker-pool    : Demonstrates worker pool pattern for controlled concurrency")
		fmt.Println("  unbounded      : Demonstrates unbounded goroutines (not recommended)")
		fmt.Println("\nOptions:")
		fmt.Println("  -profiler      : Enable pprof profiler on localhost:6060")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
