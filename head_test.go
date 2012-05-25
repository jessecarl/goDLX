package goDLX

import "testing"

// TestNew tests that creating a new Head actually does create a new Head.
// Nothing to verify other than a lack of errors.
func TestNew(t *testing.T) {
	h := New()
	if h == nil {
		t.Errorf("New(): Did not return a valid Head node")
	}
	if !assertCircleLft(h, 1) {
		t.Errorf("New(): Head not circular to the left")
	}
	if !assertCircleRgt(h, 1) {
		t.Errorf("New(): Head not circular to the right")
	}
}

func TestHeadLft(t *testing.T) {
	h := New()
	if lft := h.lft(); lft == nil {
		t.Errorf("Head.lft(): No node returned on left")
	}
}

func TestHeadSetLft(t *testing.T) {
	h := New()
	oh := new(Head)
	c := new(column)
	e := new(element)
	allowed := []node{h, c}
	forbidden := []node{oh, nil, e}
	for _, i := range allowed {
		if n, err := h.setLft(i); err != nil || n != i {
			t.Error(i)
		} else if !assertLft(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.setLft(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadRgt(t *testing.T) {
	h := New()
	if rgt := h.rgt(); rgt == nil {
		t.Errorf("Head.rgt(): No node returned on right")
	}
}

func TestHeadSetRgt(t *testing.T) {
	h := New()
	oh := new(Head)
	c := new(column)
	e := new(element)
	allowed := []node{h, c}
	forbidden := []node{oh, nil, e}
	for _, i := range allowed {
		if n, err := h.setRgt(i); err != nil || n != i {
			t.Error(i)
		} else if !assertRgt(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.setRgt(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadUp(t *testing.T) {
	h := New()
	if n := h.up(); n != h {
		t.Errorf("Head.up(): did not return self")
	}
}

func TestHeadSetUp(t *testing.T) {
	h := New()
	oh := new(Head)
	c := new(column)
	e := new(element)
	allowed := []node{}
	forbidden := []node{h, c, oh, e, nil}
	for _, i := range allowed {
		if n, err := h.setUp(i); err != nil || n != i {
			t.Error(i)
		} else if !assertUp(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.setUp(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadDn(t *testing.T) {
	h := New()
	if n := h.dn(); n != h {
		t.Errorf("Head.dn(): did not return self")
	}
}

func TestHeadSetDn(t *testing.T) {
	h := New()
	oh := new(Head)
	c := new(column)
	e := new(element)
	allowed := []node{}
	forbidden := []node{h, c, oh, e, nil}
	for _, i := range allowed {
		if n, err := h.setDn(i); err != nil || n != i {
			t.Error(i)
		} else if !assertDn(h, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := h.setDn(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestHeadCol(t *testing.T) {
	h := New()
	if n := h.col(); n != h {
		t.Errorf("Head.col(): did not return self")
	}
}

func TestHeadAddCol(t *testing.T) {
	const colCount = 100
	h := New()
	for i := 0; i < colCount; i++ {
		h.AddCol(string(i), false)
		if !assertCircleLft(h, i+2) || !assertCircleRgt(h, i+2) {
			t.Errorf("Head.AddCol(name, optional): List no longer circular")
		}
	}
}
