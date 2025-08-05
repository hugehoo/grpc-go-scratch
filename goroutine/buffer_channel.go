package goroutine

import (
	"fmt"
	"time"
)

// RunBufferChannelDemo - when producer is fast and consumer is slow,
// this pattern prevents buffer got blocked.
func RunBufferChannelDemo() {
	events := make(chan Event, 5)
	go gather(events)
	go archive(events)
	select {}
}

type Event struct {
	ID int
}

func gather(out chan<- Event) {
	for i := 0; i < 10; i++ {
		fmt.Println("produced", i)
		out <- Event{ID: i}
	}
	close(out)
}

func archive(in <-chan Event) {
	for e := range in {
		time.Sleep(300 * time.Millisecond) // slow consumer
		fmt.Println("Consumed", e.ID)
	}
}
