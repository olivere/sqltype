package sqltype

import (
	"testing"
	"time"
)

func TestScanDuration(t *testing.T) {
	var nd NullDuration
	d := 5 * time.Second
	nd.Scan(d)
	if want, have := true, nd.Valid; want != have {
		t.Errorf("want Valid=%v, have %v", want, have)
	}
	if want, have := d, nd.Duration; want != have {
		t.Errorf("want Duration=%v, have %v", want, have)
	}
}

func TestScanNilDuration(t *testing.T) {
	var nd NullDuration
	nd.Scan(nil)
	if want, have := false, nd.Valid; want != have {
		t.Errorf("want Valid=%v, have %v", want, have)
	}
}
