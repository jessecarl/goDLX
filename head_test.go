package goDLX

import (
	"strings"
	"testing"
)

type testMatrix struct {
	ColCounts map[string]int
	Cols      []string
	Rows      [][]string
	Solution  *testMatrix
}

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
		if err := h.AddCol(string(i), false); err == nil {
			t.Error(e_head_unique_cols)
		}
	}
}

func setupTestMatrix(full, solved string) (t *testMatrix) {
	t = new(testMatrix)
	t.ColCounts = make(map[string]int)
	for i, r := range strings.Split(full, "\n") {
		switch i {
		case 0:
			for _, cn := range strings.Split(r, ",") {
				t.Cols = append(t.Cols, cn)
				t.ColCounts[cn] = 0
			}
		default:
			if len(r) > 0 {
				row := []string{}
				for j := range r {
					if r[j] == '1' {
						t.ColCounts[t.Cols[j]]++
						row = append(row, t.Cols[j])
					}
				}
				t.Rows = append(t.Rows, row)
			}
		}
	}
	if len(solved) > 0 {
		t.Solution = setupTestMatrix(solved, "")
	}
	return t
}

const t1 = `A,B,C,D,E,F,G
0010110
1001001
0110010
1001000
0100001
0001101
`
const t1s = `A,B,C,D,E,F,G
1001000
0100001
0010110
`

func TestHeadAddRow(t *testing.T) {
	h := New()
	testMatrix := setupTestMatrix(t1, t1s)
	// add the columns
	for _, cn := range testMatrix.Cols {
		// construct actual
		if err := h.AddCol(cn, false); err != nil {
			t.Error(err)
		}
	}
	// add the rows
	for _, row := range testMatrix.Rows {
		if err := h.AddRow(row); err != nil {
			t.Error(err)
		}
	}
	aRows := [][]string{}
	rows := make([]bool, len(testMatrix.Rows))
	for c, ok := h.rgt().(*column); ok; c, ok = c.rgt().(*column) {
		// asserts that the column counts are correct
		if c.count() != testMatrix.ColCounts[c.label()] {
			t.Error(c)
		}
		// asserts that the rows exist
		// note that because of the way these are stored, it is
		// okay for a row to "exist" more than once
		for r, kk := c.dn().(*element); kk; r, kk = r.dn().(*element) {
			// construct the actual we can compare with
			row := []string{r.colName()}
			for e, oo := r.rgt().(*element); oo && node(e) != node(r); e, oo = e.rgt().(*element) {
				row = append(row, e.colName())
			}
			aRows = append(aRows, row)
		}
	}
	for _, act := range aRows {
		inExp := false
		for r, exp := range testMatrix.Rows {
			if len(act) != len(exp) {
				continue
			}
			thisOne := true
			for i := range act {
				hasMatch := false
				for j := range exp {
					if act[i] == exp[j] {
						hasMatch = true
						break
					}
				}
				if !hasMatch {
					thisOne = false
					break
				}
			}
			for i := range exp {
				hasMatch := false
				for j := range act {
					if exp[i] == act[j] {
						hasMatch = true
						break
					}
				}
				if !hasMatch {
					thisOne = false
					break
				}
			}
			if thisOne {
				rows[r] = true
				inExp = true
				break
			}
		}
		if !inExp {
			t.Error(act)
		}
	}
	// assert that all expected rows exist
	for r := range rows {
		if !rows[r] {
			t.Error(testMatrix.Rows[r])
		}
	}
}
