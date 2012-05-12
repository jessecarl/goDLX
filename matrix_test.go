package goDLX

import "testing"

// TestLinkLft Head -> Head
func TestLinkLftHH(t *testing.T) {
	hL := New()
	hR := New()
	if err := LinkLft(hL, hR); err == nil {
		t.Errorf("LinkLft(Head, Head): Head cannot link to another Head")
	}
	if err := LinkLft(hL, hL); err != nil {
		t.Error(err)
	}
}

// TestLinkLft Head -> Column
func TestLinkLftHC(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkLft(h, c); err != nil {
		t.Error(err)
	} else if !assertLft(c, h) || !assertRgt(h, c) {
		t.Errorf("LinkLft(Head, Column): Link not created")
	} else if !assertCircleLft(h, 2) {
		t.Errorf("LinkLft(Head, Column): List no longer circles left")
	} else if !assertCircleRgt(h, 2) {
		t.Errorf("LinkLft(Head, Column): List no longer circles right")
	}
}

// TestLinkLft Column -> Head
func TestLinkLftCH(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkLft(c, h); err != nil {
		t.Error(err)
	} else if !assertLft(h, c) || !assertRgt(c, h) {
		t.Errorf("LinkLft(Column, Head): Link not created")
	} else if !assertCircleLft(c, 2) {
		t.Errorf("LinkLft(Column, Head): List no longer circles left")
	} else if !assertCircleRgt(c, 2) {
		t.Errorf("LinkLft(Column, Head): List no longer circles right")
	}
}

// TestLinkLft Column -> Column
func TestLinkLftCC(t *testing.T) {
	lims := []int{2, 3, 4, 5}
	c := NewColumn("", false)
	for _, i := range lims {
		n := NewColumn(string(i), false)
		if err := LinkLft(c, n); err != nil {
			t.Error(err)
		} else if !assertLft(n, c) || !assertRgt(c, n) {
			t.Errorf("LinkLft(Column, Column): Link not created")
		} else if !assertCircleLft(c, i) {
			t.Errorf("LinkLft(Column, Column): List no longer circles left")
		} else if !assertCircleRgt(c, i) {
			t.Errorf("LinkLft(Column, Column): List no longer circles right")
		}
	}
}

// TestLinkRgt Head -> Head
func TestLinkRgtHH(t *testing.T) {
	hL := New()
	hR := New()
	if err := LinkRgt(hR, hL); err == nil {
		t.Errorf("LinkRgt(Head, Head): Head cannot link to another Head")
	}
	if err := LinkRgt(hR, hR); err != nil {
		t.Error(err)
	}
}

// TestLinkRgt Head -> Column
func TestLinkRgtHC(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkRgt(h, c); err != nil {
		t.Error(err)
	} else if !assertRgt(c, h) || !assertLft(h, c) {
		t.Errorf("LinkRgt(Head, Column): Link not created")
	} else if !assertCircleLft(h, 2) {
		t.Errorf("LinkRgt(Head, Column): List no longer circles left")
	} else if !assertCircleRgt(h, 2) {
		t.Errorf("LinkRgt(Head, Column): List no longer circles right")
	}
}

// TestLinkRgt Column -> Head
func TestLinkRgtCH(t *testing.T) {
	h := New()
	c := NewColumn("", false)
	if err := LinkRgt(c, h); err != nil {
		t.Error(err)
	} else if !assertRgt(h, c) || !assertLft(c, h) {
		t.Errorf("LinkRgt(Column, Head): Link not created")
	} else if !assertCircleLft(c, 2) {
		t.Errorf("LinkRgt(Column, Head): List no longer circles left")
	} else if !assertCircleRgt(c, 2) {
		t.Errorf("LinkRgt(Column, Head): List no longer circles right")
	}
}

// TestLinkRgt Column -> Column
func TestLinkRgtCC(t *testing.T) {
	lims := []int{2, 3, 4, 5}
	c := NewColumn("", false)
	for _, i := range lims {
		n := NewColumn(string(i), false)
		if err := LinkRgt(c, n); err != nil {
			t.Error(err)
		} else if !assertRgt(n, c) || !assertLft(c, n) {
			t.Errorf("LinkRgt(Column, Column): Link not created")
		} else if !assertCircleLft(c, i) {
			t.Errorf("LinkRgt(Column, Column): List no longer circles left")
		} else if !assertCircleRgt(c, i) {
			t.Errorf("LinkRgt(Column, Column): List no longer circles right")
		}
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
