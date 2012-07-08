// Copyright 2012 Jesse Allen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goDLX

import "testing"

const testName = "Test Name"

var names = []string{"Name1", ""}

func TestNewColumn(t *testing.T) {
	for _, name := range names {
		for _, b := range []bool{true, false} {
			if c := newColumn(name, b); c == nil {
				t.Errorf("newColumn(): Did not return a valid column node")
			} else if !assertCircleLft(c, 1) {
				t.Errorf("newColumn(): column not circular to left")
			} else if !assertCircleRgt(c, 1) {
				t.Errorf("newColumn(): column not circular to right")
			} else if !assertCircleUp(c, 1) {
				t.Errorf("newColumn(): column not circular up")
			} else if !assertCircleDn(c, 1) {
				t.Errorf("newColumn(): column not circular down")
			}
		}
	}
}

func TestColumnLft(t *testing.T) {
	c := newColumn(testName, false)
	if lft := c.lft(); lft == nil {
		t.Errorf("column.lft(): No node returned on left")
	}
}

func TestColumnSetLft(t *testing.T) {
	c := newColumn(testName, false)
	h := New()
	oc := new(column)
	e := new(element)
	allowed := []node{h, oc}
	forbidden := []node{c, nil, e}
	for _, i := range allowed {
		if n, err := c.setLft(i); err != nil || n != i {
			t.Error(i)
		} else if !assertLft(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.setLft(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnRgt(t *testing.T) {
	c := newColumn(testName, false)
	if lft := c.rgt(); lft == nil {
		t.Errorf("column.rgt(): No node returned on left")
	}
}

func TestColumnSetRgt(t *testing.T) {
	c := newColumn(testName, false)
	h := New()
	oc := new(column)
	e := new(element)
	allowed := []node{h, oc}
	forbidden := []node{c, nil, e}
	for _, i := range allowed {
		if n, err := c.setRgt(i); err != nil || n != i {
			t.Error(i)
		} else if !assertRgt(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.setRgt(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnUp(t *testing.T) {
	c := newColumn(testName, false)
	if lft := c.up(); lft == nil {
		t.Errorf("column.up(): No node returned on left")
	}
}

func TestColumnSetUp(t *testing.T) {
	c := newColumn(testName, false)
	h := New()
	oc := new(column)
	e := new(element)
	allowed := []node{c, e}
	forbidden := []node{h, oc, nil}
	for _, i := range allowed {
		if n, err := c.setUp(i); err != nil || n != i {
			t.Error(i)
		} else if !assertUp(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.setUp(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnDn(t *testing.T) {
	c := newColumn(testName, false)
	if lft := c.dn(); lft == nil {
		t.Errorf("column.dn(): No node returned on left")
	}
}

func TestColumnSetDn(t *testing.T) {
	c := newColumn(testName, false)
	h := New()
	oc := new(column)
	e := new(element)
	allowed := []node{c, e}
	forbidden := []node{h, oc, nil}
	for _, i := range allowed {
		if n, err := c.setDn(i); err != nil || n != i {
			t.Error(i)
		} else if !assertDn(c, i) {
			t.Error(i)
		}
	}
	for _, i := range forbidden {
		if n, err := c.setDn(i); err == nil || n != nil {
			t.Error(i)
		}
	}
}

func TestColumnCol(t *testing.T) {
	c := newColumn(testName, false)
	if n := c.col(); n != c {
		t.Errorf("column.col(): did not return self")
	}
}
