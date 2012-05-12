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

func assertCircleLft(n Node) bool {
	for i := n.Lft(); i != n; i = n.Lft() {
		if i == nil {
			return false
		}
	}
	return true
}

func assertCircleRgt(n Node) bool {
	for i := n.Rgt(); i != n; i = n.Lft() {
		if i == nil {
			return false
		}
	}
	return true
}

func assertCircleUp(n Node) bool {
	for i := n.Up(); i != n; i = n.Lft() {
		if i == nil {
			return false
		}
	}
	return true
}

func assertCircleDn(n Node) bool {
	for i := n.Dn(); i != n; i = n.Lft() {
		if i == nil {
			return false
		}
	}
	return true
}
