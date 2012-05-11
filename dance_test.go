package goDLX

import "testing"

func TestNew(t *testing.T) {
	if r := New(); r.L != r.R || r.L != r {
		t.Errorf("New Root element should link to left and right")
	}
}

func TestAddColNotRoot(t *testing.T) {
	n := new(Matrix)
	if _, err := n.AddCol("Name", false); err == nil {
		t.Errorf("Only Root elements can add new columns")
	}
}

func TestAddColSingle(t *testing.T) {
	const name = "TestName"
	r := New()
	c, err := r.AddCol(name, false)
	if err != nil {
		t.Error(err)
	}
	if c.U != c.D || c.U != c {
		t.Errorf("column not linking to self above and below")
	}
	if c.C != c {
		t.Errorf("column not linking to self as column (C)")
	}
	if c.L != r {
		t.Errorf("column not linking to root on left")
	}
	if c.R != r {
		t.Errorf("column not linking to root on right")
	}
	if r.L != c {
		t.Errorf("root not linking to column on left")
	}
	if r.R != c {
		t.Errorf("root not linking to column on right")
	}
	if c.Name != name {
		t.Errorf("Name not assigned")
	}
}

func TestAddColMult(t *testing.T) {
	const name, colCount = "TestName", 5
	var err error
	r := New()
	cols := make([](*Matrix), colCount)
	for i, v := range cols {
		v, err = r.AddCol(name, false)
		if err != nil {
			t.Error(err)
		}
		if v.R.L != v {
			t.Errorf("column not linking with column on right")
		}
		if v.L.R != v {
			t.Errorf("column not linking with column on left")
		}
		cols[i] = v
	}
	if cols[0].L != r {
		t.Errorf("column not linking to root on left")
	}
	if r.L != cols[colCount-1] {
		t.Errorf("root not linking to column on left")
	}
}
