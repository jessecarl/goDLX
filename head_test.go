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
		if err := h.AddCol(string(i), false); err == nil {
			t.Error(e_head_unique_cols)
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
	cols := []string{}
	eRows := [][]bool{}
	rows := []bool{}
	for i, r := range strings.Split(testMatrix, "\n") {
		switch i {
		case 0:
			for _, cn := range strings.Split(r, ",") {
				// construct actual
				if err := h.AddCol(cn, false); err != nil {
					t.Error(err)
				}
				// construct expected
				colCounts[cn] = 0
				cols = append(cols, cn)
			}
		default:
			if len(r) > 0 {
				if err := h.AddRow(func(i int, n string) bool {
					if r[i] == '1' {
						colCounts[n]++
						return true
					}
					return false
				}); err != nil {
					t.Error(err)
				}
				newRow := []bool{}
				// construct expected
				for _, x := range r {
					if x == '1' {
						newRow = append(newRow, true)
					} else {
						newRow = append(newRow, false)
					}
				}
				eRows = append(eRows, newRow)
				rows = append(rows, false)
			}
		}
	}
	for n := h.rgt(); n != h; n = n.rgt() {
		if c, ok := n.(*column); ok {
			// asserts that the column counts are correct
			if c.count() != colCounts[c.label()] {
				t.Error(c)
			}
			// asserts that the rows exist
			// note that because of the way these are stored, it is
			// okay for a row to "exist" more than once
			for r, kk := c.dn().(*element); kk; r, kk = r.dn().(*element) {
				// construct the actual we can compare with
				aRow := []bool{}
				for _, cn := range cols {
					if r.colName() == cn {
						aRow = append(aRow, true)
					} else {
						hasCol := false
						for e, oo := r.rgt().(*element); oo && (e != r); e = e.rgt().(*element) {
							if e.colName() == cn {
								hasCol = true
								break
							}
						}
						aRow = append(aRow, hasCol)
					}
				}
				// find actual in expected
				rowExists := false
				for i := range eRows {
					rowsMatch := true
					for j := range eRows[i] {
						if eRows[i][j] != aRow[j] {
							rowsMatch = false
							break
						}
					}
					if rowsMatch {
						rows[i] = true
						rowExists = true
						break
					}
				}
				if !rowExists {
					t.Error(aRow)
				}
			}
		}
	}
	// assert that all expected rows exist
	for r := range rows {
		if !rows[r] {
			t.Error(eRows[r])
		}
	}
}
