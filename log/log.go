/*
Package log is about logging actions
*/
package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/richardnascimento18/forge/adapter"
)

type Log struct {
	Timestamp time.Time
	Action    string
	Layer     string
	Message   string
	State     string
}

func (l *Log) WriteLogFile() error {
	logPath, err := adapter.ResolveLogDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(logPath, 0o755); err != nil {
		return fmt.Errorf("could not create path %s: %w", logPath, err)
	}

	filename := buildFilename(l.Timestamp)
	finalPath := filepath.Join(logPath, filename)

	data := l.format()
	return writeLine(finalPath, data)
}

func buildFilename(t time.Time) string {
	return t.Format("2006-01-02") + ".txt"
}

func (l *Log) format() string {
	return fmt.Sprintf(
		"%s [%s] %s: %s (state=%s)\n",
		l.Timestamp.Format(time.RFC3339),
		l.Layer,
		l.Action,
		l.Message,
		l.State,
	)
}

func writeLine(path string, line string) error {
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0o644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(line)
	return err
}
