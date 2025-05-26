package stats

import (
	"context"
	"log"
	"service/config"
	"service/pkg/ctx_log"
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	defaultCapacity = 5000
)

type (
	Writer interface {
		Insert(rows Rows) error
	}
	// TODO add gracefull shutdown
	Manager struct {
		writer        Writer
		flushInterval time.Duration

		ctx    context.Context
		cancel context.CancelFunc

		mu   sync.Mutex
		rows Rows
	}
)

func NewManager(w Writer, cfg *config.Config) *Manager {
	ctx, cancel := context.WithCancel(context.Background())

	return &Manager{
		writer:        w,
		flushInterval: cfg.Manager.FlushInterval,
		ctx:           ctx,
		cancel:        cancel,
		rows:          newRows(),
	}
}

func (m *Manager) Append(v Value) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.unsafeAppend(v)
}

func (m *Manager) AppendRows(rows Rows) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, v := range rows {
		m.unsafeAppend(v)
	}
}

func (m *Manager) unsafeAppend(v Value) {
	m.rows = append(m.rows, v)
}

func (m *Manager) StartInner() {
	l := ctx_log.LoggerFromContext(m.ctx)
	l.Info("Start stats_manager")
	go m.loop()
}

func (m *Manager) loop() {
	for {
		select {
		case <-time.After(m.flushInterval):
			m.startInserting()

		case <-m.ctx.Done():
			m.startInserting()
			return
		}
	}
}

func (m *Manager) startInserting() {

	rows := m.withdraw()
	if len(rows) == 0 {
		return
	}

	if err := m.writer.Insert(rows); err != nil {
		m.AppendRows(rows)
		return
	}

	log.Printf("Stats rows successfuly written: %d\n", len(rows))

	l := ctx_log.LoggerFromContext(m.ctx)
	l.Debug("Stats rows successfuly written",
		zap.Int("rows", len(rows)),
	)

}

func (m *Manager) withdraw() Rows {
	m.mu.Lock()
	defer m.mu.Unlock()

	rows := m.rows
	m.rows = newRows()

	return rows
}

func newRows() Rows {
	return make(Rows, 0, defaultCapacity)
}
