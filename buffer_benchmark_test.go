package main

import (
	"bytes"
	"testing"
)

type LogLine struct {
	Payload string
}

func generateLogs(n int) []LogLine {
	logs := make([]LogLine, n)
	for i := 0; i < n; i++ {
		logs[i] = LogLine{Payload: "This is a sample log line"}
	}
	return logs
}

// ✅ 버퍼 재사용 버전
func BenchmarkReuseBuffer(b *testing.B) {
	logLines := generateLogs(100_000)
	buf := bytes.NewBuffer(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range logLines {
			buf.Reset()
			buf.WriteString(line.Payload)
			_ = buf.Bytes()
		}
	}
}

// 🚫 버퍼 새로 생성 버전
func BenchmarkNewBufferEachTime(b *testing.B) {
	logLines := generateLogs(100_000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range logLines {
			buf := bytes.NewBuffer(nil)
			buf.WriteString(line.Payload)
			_ = buf.Bytes()
		}
	}
}
