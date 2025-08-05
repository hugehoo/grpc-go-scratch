package goroutine

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// InitProfiler starts the pprof server for profiling
func InitProfiler() {
	fmt.Println("init profiler")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

// Task
// 1. Worker pool - limit the number of goroutines
// fix the number of workers for preventing create goroutines limitless
type Task struct {
	ID int
}

func process(t Task) {
	fmt.Println("Processing task:", t.ID)
}

// RunWorkerPoolDemo demonstrates the worker pool pattern
func RunWorkerPoolDemo() {
	WorkerPools()
}

// WorkerPools demonstrates the worker pool pattern implementation
func WorkerPools() {
	jobs := make(chan Task)
	var wg sync.WaitGroup

	// worker 4 setting
	for w := 0; w < 4; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range jobs {
				fmt.Printf("[Worker %d]", id)
				process(task)
			}
		}(w)
	}

	// task 10
	for i := 0; i < 10; i++ {
		jobs <- Task{ID: i}
	}
	close(jobs)
	wg.Wait()

	fmt.Println("Done. Press Ctrl+C to quit.")
	select {}
}

// RunUnboundedGoroutinesDemo - demonstrates unbounded goroutines (not recommended)
func RunUnboundedGoroutinesDemo() {
	UnboundedGoroutines()
}

// UnboundedGoroutines - demonstrates unbounded goroutines (not recommended)
func UnboundedGoroutines() {
	var wg sync.WaitGroup
	tasks := make([]Task, 0)

	// 작업 10000개 생성
	for i := 0; i < 10000; i++ {
		tasks = append(tasks, Task{ID: i})
	}

	// 작업마다 고루틴 생성 (무제한)
	for _, t := range tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			process(task)
		}(t)
	}

	wg.Wait()
	fmt.Println("Done. Press Ctrl+C to quit.")
	select {}
}
