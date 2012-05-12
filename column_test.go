package goDLX

import "testing"

const testName = "Test Name"

var names = []string{"Name1", ""}

func TestNewColumn(t *testing.T) {
	for _, name := range names {
		for _, b := range []bool{true, false} {
			if c := NewColumn(name, b); c == nil {
				t.Errorf("NewColumn(): Did not return a valid Column Node")
			} else if !assertCircleLft(c) {
				t.Errorf("NewColumn(): Column not circular to left")
			} else if !assertCircleRgt(c) {
				t.Errorf("NewColumn(): Column not circular to right")
			} else if !assertCircleUp(c) {
				t.Errorf("NewColumn(): Column not circular up")
			} else if !assertCircleDn(c) {
				t.Errorf("NewColumn(): Column not circular down")
			}
		}
	}
}

func TestColumnLft(t *testing.T) {
	c := NewColumn(testName, false)
	if lft := c.Lft(); lft == nil {
		t.Errorf("Column.Lft(): No Node returned on left")
	}
}

func TestColumnSetLft(t *testing.T) {
	c := NewColumn(testName, false)
	h := New()
	oc := new(Column)
	allowed := []Node{h, oc}
	forbidden := []Node{c, nil}
	for _, i := range allowed {
		if n, err := c.SetLft(i); err != nil || n != i {
			t.Error(i)
		} else if !assertLft(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.SetLft(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnRgt(t *testing.T) {
	c := NewColumn(testName, false)
	if lft := c.Rgt(); lft == nil {
		t.Errorf("Column.Rgt(): No Node returned on left")
	}
}

func TestColumnSetRgt(t *testing.T) {
	c := NewColumn(testName, false)
	h := New()
	oc := new(Column)
	allowed := []Node{h, oc}
	forbidden := []Node{c, nil}
	for _, i := range allowed {
		if n, err := c.SetRgt(i); err != nil || n != i {
			t.Error(i)
		} else if !assertRgt(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.SetRgt(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnUp(t *testing.T) {
	c := NewColumn(testName, false)
	if lft := c.Up(); lft == nil {
		t.Errorf("Column.Up(): No Node returned on left")
	}
}

func TestColumnSetUp(t *testing.T) {
	c := NewColumn(testName, false)
	h := New()
	oc := new(Column)
	allowed := []Node{c}
	forbidden := []Node{h, oc, nil}
	for _, i := range allowed {
		if n, err := c.SetUp(i); err != nil || n != i {
			t.Error(i)
		} else if !assertUp(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.SetUp(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnDn(t *testing.T) {
	c := NewColumn(testName, false)
	if lft := c.Dn(); lft == nil {
		t.Errorf("Column.Dn(): No Node returned on left")
	}
}

func TestColumnSetDn(t *testing.T) {
	c := NewColumn(testName, false)
	h := New()
	oc := new(Column)
	allowed := []Node{c}
	forbidden := []Node{h, oc, nil}
	for _, i := range allowed {
		if n, err := c.SetDn(i); err != nil || n != i {
			t.Error(i)
		} else if !assertDn(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.SetDn(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnCol(t *testing.T) {
	c := NewColumn(testName, false)
	if n := c.Col(); n != c {
		t.Errorf("Column.Col(): did not return self")
	}
}
