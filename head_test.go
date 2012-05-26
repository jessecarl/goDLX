package goDLX

import (
	"strings"
	"testing"
)

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

const testMatrix = `A,B,C,D,E,F,G
0010110
1001001
0110010
1001000
0100001
0001101
`
const solvedMatrix = `A,B,C,D,E,F,G
1001000
0100001
0010110
`

func TestHeadAddRow(t *testing.T) {
	h := New()
	// add the columns
	colCounts := make(map[string]int)
	colCounts["A"] = 2
	colCounts["B"] = 2
	colCounts["C"] = 2
	colCounts["D"] = 3
	colCounts["E"] = 2
	colCounts["F"] = 2
	colCounts["G"] = 3
	for i, r := range strings.Split(testMatrix, "\n") {
		switch i {
		case 0:
			for _, cn := range strings.Split(r, ",") {
				if err := h.AddCol(cn, false); err != nil {
					t.Error(err)
				}
			}
		default:
			if len(r) > 0 {
				if err := h.AddRow(func(i int, n string) bool {
					if r[i] == '1' {
						return true
					}
					return false
				}); err != nil {
					t.Error(err)
				}
			}
		}
	}
	// asserts that the column counts are correct
	// not perfect, but a better test than many
	for n := h.rgt(); n != h; n = n.rgt() {
		if c, ok := n.(*column); ok {
			if c.count() != colCounts[c.label()] {
				t.Error(c)
			}
		}
	}
}
