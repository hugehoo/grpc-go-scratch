package goroutine

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

func transform(chunk []byte) {
	fmt.Printf("Chunk: %s\n", chunk)
}

func StreamChunks() {
	data := bytes.NewReader([]byte("HelloWorldThisIsAStreamOfData"))
	chuck := make([]byte, 8)

	for {
		n, err := io.ReadFull(data, chuck)
		if n > 0 {
			transform(chuck[:n])
		}
		if err == io.EOF || errors.Is(err, io.ErrUnexpectedEOF) {
			break
		}
	}
}
