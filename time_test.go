package sqltype

import (
	"testing"
	"time"
)

func TestScanTime(t *testing.T) {
	var nt NullTime
	ts := time.Now()
	nt.Scan(ts)
	if want, have := true, nt.Valid; want != have {
		t.Errorf("want Valid=%v, have %v", want, have)
	}
	if want, have := ts, nt.Time; want != have {
		t.Errorf("want Time=%v, have %v", want, have)
	}
}

func TestScanNilTime(t *testing.T) {
	var nt NullTime
	nt.Scan(nil)
	if want, have := false, nt.Valid; want != have {
		t.Errorf("want Valid=%v, have %v", want, have)
	}
}
