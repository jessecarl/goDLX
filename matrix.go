package goDLX

// The Node interface navigates a sparse matrix.
// Any type implementing a 2-dimensional doubly-linked list satisfies this.
// (despite the convention being that interfaces have the -er suffix, this special case
// can't do that without being entirely unclear)
type Node interface {
	Lft() Node
	SetLft(Node) (Node, error)
	Rgt() Node
	SetRgt(Node) (Node, error)
	Up() Node
	SetUp(Node) (Node, error)
	Dn() Node
	SetDn(Node) (Node, error)
	Col() Node
}
