package goDLX

import "testing"

// TestLinkHorz Head -> Head
func TestLinkHorzHH(t *testing.T) {
	hL := New()
	hR := New()
	if err := linkHorz(hL, hR); err == nil {
		t.Errorf("linkHorz(Head, Head): Head cannot link to another Head")
	}
	if err := linkHorz(hL, hL); err != nil {
		t.Error(err)
	}
}

// TestLinkHorz Head -> column
func TestLinkHorzHC(t *testing.T) {
	h := New()
	c := newColumn("", false)
	if err := linkHorz(h, c); err != nil {
		t.Error(err)
	} else if !assertLft(c, h) || !assertRgt(h, c) {
		t.Errorf("linkHorz(Head, column): Link not created")
	} else if !assertCircleLft(h, 2) {
		t.Errorf("linkHorz(Head, column): List no longer circles left")
	} else if !assertCircleRgt(h, 2) {
		t.Errorf("linkHorz(Head, column): List no longer circles right")
	}
}

// TestLinkHorz column -> Head
func TestLinkHorzCH(t *testing.T) {
	h := New()
	c := newColumn("", false)
	if err := linkHorz(c, h); err != nil {
		t.Error(err)
	} else if !assertLft(h, c) || !assertRgt(c, h) {
		t.Errorf("linkHorz(column, Head): Link not created")
	} else if !assertCircleLft(c, 2) {
		t.Errorf("linkHorz(column, Head): List no longer circles left")
	} else if !assertCircleRgt(c, 2) {
		t.Errorf("linkHorz(column, Head): List no longer circles right")
	}
}

// TestLinkHorz column -> column
func TestLinkHorzCC(t *testing.T) {
	lims := []int{2, 3, 4, 5}
	c := newColumn("", false)
	if err := linkHorz(c, c); err != nil {
		t.Error(err)
	} else if !assertLft(c, c) || !assertRgt(c, c) {
		t.Errorf("linkHorz(column, column): Link not created")
	} else if !assertCircleLft(c, 1) {
		t.Errorf("linkHorz(column, column): List no longer circles left")
	} else if !assertCircleRgt(c, 1) {
		t.Errorf("linkHorz(column, column): List no longer circles right")
	}
	for _, i := range lims {
		n := newColumn(string(i), false)
		if err := linkHorz(c, n); err != nil {
			t.Error(err)
		} else if !assertLft(n, c) || !assertRgt(c, n) {
			t.Errorf("linkHorz(column, column): Link not created")
		} else if !assertCircleLft(c, i) {
			t.Errorf("linkHorz(column, column): List no longer circles left")
		} else if !assertCircleRgt(c, i) {
			t.Errorf("linkHorz(column, column): List no longer circles right")
		}
	}
}

// TestLinkVert Head -> Head
func TestLinkVertHH(t *testing.T) {
	hU := New()
	hD := New()
	if err := linkVert(hU, hD); err == nil {
		t.Errorf("linkVert(Head, Head): Head cannot link vertically")
	}
}

// TestLinkVert Head -> column
func TestLinkVertHC(t *testing.T) {
	h := New()
	c := newColumn("", false)
	if err := linkVert(h, c); err == nil {
		t.Errorf("linkVert(Head, column): Head cannot link vertically")
	}
}

// TestLinkVert column -> Head
func TestLinkVertCH(t *testing.T) {
	h := New()
	c := newColumn("", false)
	if err := linkVert(c, h); err == nil {
		t.Errorf("linkVert(column, Head): Head cannot link vertically")
	}
}

// TestLinkVert column -> column
func TestLinkVertCC(t *testing.T) {
	c := newColumn("", false)
	co := newColumn("", false)
	if err := linkVert(c, c); err != nil {
		t.Error(err)
	} else if !assertUp(c, c) || !assertDn(c, c) {
		t.Errorf("linkVert(column, column): Link not created")
	} else if !assertCircleUp(c, 1) {
		t.Errorf("linkVert(column, column): List no longer circles up")
	} else if !assertCircleDn(c, 1) {
		t.Errorf("linkVert(column, column): List no longer circles down")
	}
	if err := linkVert(c, co); err == nil {
		t.Errorf("linkVert(column, column): column cannot link vertically to another column")
	}
}

// Assert functions
func assertLft(check node, expected node) bool {
	return check.lft() == expected
}

func assertRgt(check node, expected node) bool {
	return check.rgt() == expected
}

func assertUp(check node, expected node) bool {
	return check.up() == expected
}

func assertDn(check node, expected node) bool {
	return check.dn() == expected
}

func assertCircleLft(n node, lim int) bool {
	count := 0
	for i := n.lft(); i != n; i = i.lft() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}

func assertCircleRgt(n node, lim int) bool {
	count := 0
	for i := n.rgt(); i != n; i = i.rgt() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}

func assertCircleUp(n node, lim int) bool {
	count := 0
	for i := n.up(); i != n; i = i.up() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}

func assertCircleDn(n node, lim int) bool {
	count := 0
	for i := n.dn(); i != n; i = i.dn() {
		count++
		if i == nil || count >= lim {
			return false
		}
	}
	return true
}
