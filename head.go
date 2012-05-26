package goDLX

import "errors"

const (
	e_head_set_horz = "Only column of Self can be left or right of Head"
	e_head_set_vert = "No nodes can be above or below Head"
	e_head_locked   = "All Columns must be added before any Rows"
	e_head_row_fail = "Failed to add row to matrix"
)

// Head nodes are the master column headers.
// These nodes form the heart of each sparse matrix
type Head struct {
	left, right node
	locked      bool // locks the matrix from adding new columns once rows are added
}

// New starts a new sparse matrix by creating the Head node.
func New() *Head {
	h := new(Head)
	h.initHead()
	return h
}

func (h *Head) initHead() {
	h.left = h
	h.right = h
}

////
// Satisfy the node interface
///

// lft returns the node to the left of the Head
func (h *Head) lft() node {
	return h.left
}

// setLft sets a node to the left of the Head node.
func (h *Head) setLft(n node) (node, error) {
	switch n.(type) {
	case *Head:
		// only invalid if it is a different Head
		if h != n {
			return nil, errors.New(e_head_set_horz)
		}
	case *column:
		// do nothing, we're good
	default:
		return nil, errors.New(e_head_set_horz)
	}
	h.left = n
	return n, nil
}

// rgt returns the node to the right of the Head
func (h *Head) rgt() node {
	return h.right
}

// setRgt sets the node to the right of the Head
func (h *Head) setRgt(n node) (node, error) {
	switch n.(type) {
	case *Head:
		// only invalid if it is a different Head
		if h != n {
			return nil, errors.New(e_head_set_horz)
		}
	case *column:
		// do nothing, we're good
	default:
		return nil, errors.New(e_head_set_horz)
	}
	h.right = n
	return n, nil
}

// up returns the Head itself as Head nodes are only horizontally linked
func (h *Head) up() node {
	return h
}

// setUp always errors because Head nodes are horizontally linked only
func (h *Head) setUp(n node) (node, error) {
	return nil, errors.New(e_head_set_vert)
}

// dn returns the Head itself as Head nodes are only horizontally linked
func (h *Head) dn() node {
	return h
}

// setDn always errors because Head nodes are horizontally linked only
func (h *Head) setDn(n node) (node, error) {
	return nil, errors.New(e_head_set_vert)
}

// col returns the Head itself as it is the header in the column list
func (h *Head) col() node {
	return h
}

////
// Populating the Matrix
////

// AddCol adds a Column to the matrix.
// Each column has a name for reference and can be set to optional.
// Where required columns have one and only one valid row, optional
// columns have zero or one valid row.
func (h *Head) AddCol(name string, optional bool) error {
	if h.locked {
		return errors.New(e_head_locked)
	}
	c := newColumn(name, optional)
	err := linkHorz(c, h)
	return err
}

// AddRow adds a row to the matrix
// Each row is populated by adding links for every true returned
// from the anonymous function whose arguments are the column
// name and index. Once rows are added, the matrix is locked
// and no further columns may be added.
func (h *Head) AddRow(f func(int, string) bool) error {
	h.locked = true // lock this matrix
	// loop through the columns to check against
	i := 0
	var e *element
	for n := h.rgt(); n != h; n = n.rgt() {
		if c, ok := n.(*column); ok {
			if f(i, c.name) {
				e = newElement(e, c)
			}
		} else {
			return errors.New(e_head_row_fail)
		}
		i++
	}
	if e == nil {
		return errors.New(e_head_row_fail)
	}
	return nil
}
