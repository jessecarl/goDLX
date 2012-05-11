package goDLX

import "testing"

func TestNew(t *testing.T) {
	if r := New(); r.L != r.R || r.L != r {
		t.Errorf("New Root element should link to left and right")
	}
}
