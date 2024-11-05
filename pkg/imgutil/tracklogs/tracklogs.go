package tracklogs

import (
	"fmt"
	"os"
	"time"
)

const logsPath = "/tmp/nerdctl-tracklogs"

type TrackLogs struct {
	file *os.File

	beginTime time.Time
	operation string
}

func New() (*TrackLogs, error) {
	file, err := os.OpenFile(logsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return &TrackLogs{file: file}, nil
}

func (t *TrackLogs) Begin(operation string) {
	t.beginTime = time.Now()
	t.operation = operation
	t.WriteLogs(fmt.Sprintf("[TRACK] Begin %s\n", operation))
}

func (t *TrackLogs) End() {
	elapsed := time.Since(t.beginTime)
	t.WriteLogs(fmt.Sprintf("[TRACK] End %s, elapsed: %s\n", t.operation, elapsed))
}

func (t *TrackLogs) WriteLogs(msg string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	_, err := t.file.WriteString(timeStr + " " + msg)
	if err != nil {
		panic("Write tracklogs failed: " + err.Error())
	}
}

func (t *TrackLogs) Close() {
	err := t.file.Close()
	if err != nil {
		panic("Close tracklogs failed: " + err.Error())
	}
}
