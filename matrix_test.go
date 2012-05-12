package goDLX

import "testing"

// TestLinkHorz Head -> Head
func TestLinkHorzHH(t *testing.T) {
	hL := New()
	hR := New()
	if err := LinkHorz(hL, hR); err == nil {
		t.Errorf("LinkHorz(Head, Head): Head cannot link to another Head")
	}
	if err := LinkHorz(hL, hL); err != nil {
		t.Error(err)
	}
}

// TestLinkHorz Head -> Column
func TestLinkHorzHC(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkHorz(h, c); err != nil {
		t.Error(err)
	} else if !assertLft(c, h) || !assertRgt(h, c) {
		t.Errorf("LinkHorz(Head, Column): Link not created")
	} else if !assertCircleLft(h, 2) {
		t.Errorf("LinkHorz(Head, Column): List no longer circles left")
	} else if !assertCircleRgt(h, 2) {
		t.Errorf("LinkHorz(Head, Column): List no longer circles right")
	}
}

// TestLinkHorz Column -> Head
func TestLinkHorzCH(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkHorz(c, h); err != nil {
		t.Error(err)
	} else if !assertLft(h, c) || !assertRgt(c, h) {
		t.Errorf("LinkHorz(Column, Head): Link not created")
	} else if !assertCircleLft(c, 2) {
		t.Errorf("LinkHorz(Column, Head): List no longer circles left")
	} else if !assertCircleRgt(c, 2) {
		t.Errorf("LinkHorz(Column, Head): List no longer circles right")
	}
}

// TestLinkHorz Column -> Column
func TestLinkHorzCC(t *testing.T) {
	lims := []int{2, 3, 4, 5}
	c := NewColumn("", false)
	if err := LinkHorz(c, c); err != nil {
		t.Error(err)
	} else if !assertLft(c, c) || !assertRgt(c, c) {
		t.Errorf("LinkHorz(Column, Column): Link not created")
	} else if !assertCircleLft(c, 1) {
		t.Errorf("LinkHorz(Column, Column): List no longer circles left")
	} else if !assertCircleRgt(c, 1) {
		t.Errorf("LinkHorz(Column, Column): List no longer circles right")
	}
	for _, i := range lims {
		n := NewColumn(string(i), false)
		if err := LinkHorz(c, n); err != nil {
			t.Error(err)
		} else if !assertLft(n, c) || !assertRgt(c, n) {
			t.Errorf("LinkHorz(Column, Column): Link not created")
		} else if !assertCircleLft(c, i) {
			t.Errorf("LinkHorz(Column, Column): List no longer circles left")
		} else if !assertCircleRgt(c, i) {
			t.Errorf("LinkHorz(Column, Column): List no longer circles right")
		}
	}
}

// TestLinkVert Head -> Head
func TestLinkVertHH(t *testing.T) {
	hU := New()
	hD := New()
	if err := LinkVert(hU, hD); err == nil {
		t.Errorf("LinkVert(Head, Head): Head cannot link vertically")
	}
}

// TestLinkVert Head -> Column
func TestLinkVertHC(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkVert(h, c); err == nil {
		t.Errorf("LinkVert(Head, Column): Head cannot link vertically")
	}
}

// TestLinkVert Column -> Head
func TestLinkVertCH(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkVert(c, h); err == nil {
		t.Errorf("LinkVert(Column, Head): Head cannot link vertically")
	}
}

// TestLinkVert Column -> Column
func TestLinkVertCC(t *testing.T) {
	c := NewColumn("", false)
	co := NewColumn("", false)
	if err := LinkVert(c, c); err != nil {
		t.Error(err)
	} else if !assertUp(c, c) || !assertDn(c, c) {
		t.Errorf("LinkVert(Column, Column): Link not created")
	} else if !assertCircleUp(c, 1) {
		t.Errorf("LinkVert(Column, Column): List no longer circles up")
	} else if !assertCircleDn(c, 1) {
		t.Errorf("LinkVert(Column, Column): List no longer circles down")
	}
	if err := LinkVert(c, co); err == nil {
		t.Errorf("LinkVert(Column, Column): Column cannot link vertically to another Column")
	}
}

// Assert functions
func assertLft(check Node, expected Node) bool {
	return check.Lft() == expected
}

func assertRgt(check Node, expected Node) bool {
	return check.Rgt() == expected
}

func assertUp(check Node, expected Node) bool {
	return check.Up() == expected
}

func assertDn(check Node, expected Node) bool {
	return check.Dn() == expected
}

func assertCircleLft(n Node, lim int) bool {
	count := 0
	for i := n.Lft(); i != n; i = i.Lft() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}

func assertCircleRgt(n Node, lim int) bool {
	count := 0
	for i := n.Rgt(); i != n; i = i.Rgt() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}

func assertCircleUp(n Node, lim int) bool {
	count := 0
	for i := n.Up(); i != n; i = i.Up() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}

func assertCircleDn(n Node, lim int) bool {
	count := 0
	for i := n.Dn(); i != n; i = i.Dn() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}
