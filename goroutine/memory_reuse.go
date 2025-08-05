package goroutine

import (
	"bytes"
	"fmt"
)

type LogLine struct {
	Payload string
}

func ship(data []byte) {
	fmt.Printf("shipping: %s\n", string(data))
}

func MemoryReuse() {

	logLines := []LogLine{
		{"Hello"},
		{"World"},
		{"Gophers"},
	}

	buf := bytes.NewBuffer(nil)
	for _, l := range logLines {
		buf.Reset()
		buf.WriteString(l.Payload)
		ship(buf.Bytes())
	}
}
