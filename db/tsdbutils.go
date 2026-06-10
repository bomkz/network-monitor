package db

import (
	"context"
	"sync"
	"time"

	"github.com/prometheus/prometheus/tsdb"
)

type TSDBWriter struct {
	db       *tsdb.DB
	pointsCh chan TSDBPoint
	errorsCh chan error
	wg       sync.WaitGroup
}

func NewTSDBWriter(db *tsdb.DB, bufferSize int) *TSDBWriter {
	w := &TSDBWriter{
		db:       db,
		pointsCh: make(chan TSDBPoint, bufferSize),
		errorsCh: make(chan error, 16),
	}
	w.wg.Add(1)
	go w.run()
	return w
}

// Non-blocking write from any goroutine
func (w *TSDBWriter) Write(p TSDBPoint) {
	w.pointsCh <- p
}

// Background goroutine — owns the Appender exclusively
func (w *TSDBWriter) run() {
	defer w.wg.Done()

	// Batch flush every 5 seconds or every 100 points
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	var batch []TSDBPoint

	for {
		select {
		case p, ok := <-w.pointsCh:
			if !ok {
				// Channel closed, flush remaining
				w.flush(batch)
				return
			}
			batch = append(batch, p)
			if len(batch) >= 100 {
				w.flush(batch)
				batch = batch[:0]
			}

		case <-ticker.C:
			if len(batch) > 0 {
				w.flush(batch)
				batch = batch[:0]
			}
		}
	}
}

func (w *TSDBWriter) flush(batch []TSDBPoint) {
	if len(batch) == 0 {
		return
	}

	app := w.db.Appender(context.Background())

	for _, p := range batch {
		if err := WritePoint(app, p); err != nil {
			_ = app.Rollback()
			w.errorsCh <- err
			return
		}
	}

	if err := app.Commit(); err != nil {
		w.errorsCh <- err
	}
}

func (w *TSDBWriter) Close() {
	close(w.pointsCh)
	w.wg.Wait()
}
