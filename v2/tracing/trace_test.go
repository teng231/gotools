package tracing

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	trace := &Timer{
		StartAt:      time.Now(),
		Label:        "PayVoucher",
		IsExportable: true, //
		Verbose:      false,
		ExportData:   make(map[string]string),
	}
	defer func(trace *Timer) {
		out, _ := json.MarshalIndent(trace.ExportData, "", " ")
		log.Print(string(out))
	}(trace)
	d := trace.PinNext("s1")
	log.Print("s1 ", d)

	time.Sleep(120 * time.Millisecond)
	d = trace.PinNext("s2")
	log.Print("s2 ", d)
	time.Sleep(200 * time.Millisecond)
	trace.PinNext("s3")
	log.Print("s3 ", d)

	time.Sleep(400 * time.Millisecond)
	d = trace.PinNext("s4")
	log.Print("s4 ", d)
}
