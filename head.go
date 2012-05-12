package goDLX

import "errors"

const (
	ERROR_HEAD_SET_HORZ = "Only Column of Self can be left or right of Head"
	ERROR_HEAD_SET_VERT = "No Nodes can be above or below Head"
)

// Head nodes are the master column headers.
// These nodes form the heart of each sparse matrix
type Head struct {
	lft, rgt Node
	locked   bool // locks the matrix from adding new columns once rows are added
}

// New starts a new sparse matrix by creating the Head node.
func New() *Head {
	h := new(Head)
	h.initHead()
	return h
}

func (h *Head) initHead() {
	h.lft = h
	h.rgt = h
}

////
// Satisfy the Node interface
///

// Lft returns the node to the left of the Head
func (h *Head) Lft() Node {
	return h.lft
}

// SetLft sets a node to the left of the Head node.
func (h *Head) SetLft(n Node) (Node, error) {
	switch n.(type) {
	case *Head:
		// only invalid if it is a different Head
		if h != n {
			return nil, errors.New(ERROR_HEAD_SET_HORZ)
		}
	default:
		return nil, errors.New(ERROR_HEAD_SET_HORZ)
	}
	h.lft = n
	return n, nil
}

// Rgt returns the node to the right of the Head
func (h *Head) Rgt() Node {
	return h.rgt
}

// SetRgt sets the node to the right of the Head
func (h *Head) SetRgt(n Node) (Node, error) {
	switch n.(type) {
	case *Head:
		// only invalid if it is a different Head
		if h != n {
			return nil, errors.New(ERROR_HEAD_SET_HORZ)
		}
	default:
		return nil, errors.New(ERROR_HEAD_SET_HORZ)
	}
	h.rgt = n
	return n, nil
}

// Up returns the Head itself as Head nodes are only horizontally linked
func (h *Head) Up() Node {
	return h
}

// SetUp always errors because Head Nodes are horizontally linked only
func (h *Head) SetUp(n Node) (Node, error) {
	return nil, errors.New(ERROR_HEAD_SET_VERT)
}

// Dn returns the Head itself as Head nodes are only horizontally linked
func (h *Head) Dn() Node {
	return h
}

// SetDn always errors because Head Nodes are horizontally linked only
func (h *Head) SetDn(n Node) (Node, error) {
	return nil, errors.New(ERROR_HEAD_SET_VERT)
}

// Col returns the Head itself as it is the header in the column list
func (h *Head) Col() Node {
	return h
}
