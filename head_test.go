package goDLX

import "testing"

// TestNew tests that creating a new Head actually does create a new Head.
// Nothing to verify other than a lack of errors.
func TestNew(t *testing.T) {
	h := New()
	if h == nil {
		t.Errorf("New(): Did not return a valid Head node")
	}
	if !assertCircleLft(h) {
		t.Errorf("New(): Head not circular to the left")
	}
	if !assertCircleRgt(h) {
		t.Errorf("New(): Head not circular to the right")
	}
}

func TestHeadLft(t *testing.T) {
	h := New()
	if lft := h.Lft(); lft == nil {
		t.Errorf("Head.Lft(): No Node returned on left")
	}
}

func TestHeadSetLft(t *testing.T) {
	h := New()
	oh := new(Head)
	allowed := []Node{h}
	forbidden := []Node{oh, nil}
	for _, i := range allowed {
		if n, err := h.SetLft(i); err != nil || n != i {
			t.Error(i)
		} else if !assertLft(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.SetLft(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadRgt(t *testing.T) {
	h := New()
	if rgt := h.Rgt(); rgt == nil {
		t.Errorf("Head.Rgt(): No Node returned on right")
	}
}

func TestHeadSetRgt(t *testing.T) {
	h := New()
	oh := new(Head)
	allowed := []Node{h}
	forbidden := []Node{oh, nil}
	for _, i := range allowed {
		if n, err := h.SetRgt(i); err != nil || n != i {
			t.Error(i)
		} else if !assertRgt(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.SetRgt(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadUp(t *testing.T) {
	h := New()
	if n := h.Up(); n != h {
		t.Errorf("Head.Up(): did not return self")
	}
}

func TestHeadSetUp(t *testing.T) {
	h := New()
	oh := new(Head)
	allowed := []Node{}
	forbidden := []Node{h, oh, nil}
	for _, i := range allowed {
		if n, err := h.SetUp(i); err != nil || n != i {
			t.Error(i)
		} else if !assertUp(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.SetUp(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadDn(t *testing.T) {
	h := New()
	if n := h.Dn(); n != h {
		t.Errorf("Head.Dn(): did not return self")
	}
}

func TestHeadSetDn(t *testing.T) {
	h := New()
	oh := new(Head)
	allowed := []Node{}
	forbidden := []Node{h, oh, nil}
	for _, i := range allowed {
		if n, err := h.SetDn(i); err != nil || n != i {
			t.Error(i)
		} else if !assertDn(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.SetDn(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadCol(t *testing.T) {
	h := New()
	if n := h.Col(); n != h {
		t.Errorf("Head.Col(): did not return self")
	}
}
