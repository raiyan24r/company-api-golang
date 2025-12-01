package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// dailyFileWriter implements zapcore.WriteSyncer and rotates log files by date.
// Each day it writes to logs/app-YYYY-MM-DD.log creating the file if needed.
type dailyFileWriter struct {
    mu       sync.Mutex
    dir      string
    prefix   string
    curDate  string
    file     *os.File
}

// newDailyFileWriter constructs a writer in the provided directory with the given prefix.
func newDailyFileWriter(dir, prefix string) (*dailyFileWriter, error) {
    if dir == "" {
        return nil, errors.New("directory required")
    }
    if err := os.MkdirAll(dir, 0755); err != nil {
        return nil, fmt.Errorf("create logs dir: %w", err)
    }
    w := &dailyFileWriter{dir: dir, prefix: prefix}
    if err := w.rotateIfNeeded(); err != nil {
        return nil, err
    }
    return w, nil
}

// rotateIfNeeded checks the current date and opens a new file if the day changed.
func (w *dailyFileWriter) rotateIfNeeded() error {
    today := time.Now().Format("2006-01-02")
    if w.file != nil && today == w.curDate {
        return nil // still same day
    }
    // Close old file if present
    if w.file != nil {
        _ = w.file.Close()
        w.file = nil
    }
    w.curDate = today
    name := fmt.Sprintf("%s-%s.log", w.prefix, today)
    fullPath := filepath.Join(w.dir, name)
    f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("open log file: %w", err)
    }
    w.file = f
    return nil
}

// Write implements io.Writer; it rotates before writing if the day changed.
func (w *dailyFileWriter) Write(p []byte) (int, error) {
    w.mu.Lock()
    defer w.mu.Unlock()
    if err := w.rotateIfNeeded(); err != nil {
        return 0, err
    }
    return w.file.Write(p)
}

// Sync flushes file contents to disk.
func (w *dailyFileWriter) Sync() error {
    w.mu.Lock()
    defer w.mu.Unlock()
    if w.file == nil {
        return nil
    }
    return w.file.Sync()
}

// New creates a zap.Logger that logs JSON lines to stderr and a daily rotating file.
func New() (*zap.Logger, error) {
    writer, err := newDailyFileWriter("logs", "app")
    if err != nil {
        return nil, err
    }

    encCfg := zap.NewProductionEncoderConfig()
    encCfg.TimeKey = "ts"
    encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
    encoder := zapcore.NewJSONEncoder(encCfg)

    // File core
    fileCore := zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.InfoLevel)
    // Stderr core
    // stderrCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), zap.InfoLevel)

    // Tee both cores so each log goes to file and stderr
    core := zapcore.NewTee(fileCore /*, stderrCore*/)
    logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
    return logger, nil
}
