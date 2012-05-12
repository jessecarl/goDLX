package goDLX

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
