package tracing

import (
	"log"
	"strconv"
	"time"
)

type Timer struct {
	index   int
	StartAt time.Time

	Label        string
	IsExportable bool
	ExportData   map[string]string
	Verbose      bool
}

func (t *Timer) Start() {
	t.StartAt = time.Now()
	t.index = 0
}

func (t *Timer) Pin(description string) time.Duration {
	d := time.Since(t.StartAt)
	if t.Verbose {
		log.Printf("%s[%d]: %s %s", t.Label, t.index, description, d)
	}

	if t.IsExportable {
		t.ExportData[t.Label+strconv.Itoa(t.index)] = description + ":" + d.String()
	}
	t.index++
	return d
}

func (t *Timer) PinNext(description string) time.Duration {
	d := time.Since(t.StartAt)

	if t.Verbose {
		log.Printf("%s[%d]: %s %s", t.Label, t.index, description, d)
	}
	if t.IsExportable {
		t.ExportData[t.Label+strconv.Itoa(t.index)] = description + ":" + d.String()
	}
	t.StartAt = time.Now()
	t.index++
	return d
}

func (t *Timer) Export() map[string]string {
	return t.ExportData
}
