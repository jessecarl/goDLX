package goDLX

import "errors"

const (
	e_element_set_horz = "Only elements can be left or right of element"
	e_element_set_vert = "Only element or column can be above or below element"
)

type element struct {
	left, right, upup, down, colcol node
}

func newElement(r *element, c *column) *element {
	e := new(element)
	e.init(r, c)
	return e
}

func (e *element) init(r *element, c *column) {
	e.left = e
	e.right = e
	e.upup = e
	e.down = e
	if r != nil {
		// to the right
		linkHorz(e, r.rgt())
	}
	if c != nil {
		// above (at the bottom)
		linkVert(e, c)
		e.colcol = node(c)
	}
}

func (e *element) lft() node {
	return e.left
}

func (e *element) rgt() node {
	return e.right
}

func (e *element) up() node {
	return e.upup
}

func (e *element) dn() node {
	return e.down
}

func (e *element) col() node {
	return e.colcol
}

func (e *element) setLft(n node) (node, error) {
	switch n.(type) {
	case *element:
		e.left = n
	default:
		return nil, errors.New(e_element_set_horz)
	}
	return n, nil
}

func (e *element) setRgt(n node) (node, error) {
	switch n.(type) {
	case *element:
		e.right = n
	default:
		return nil, errors.New(e_element_set_horz)
	}
	return n, nil
}

func (e *element) setUp(n node) (node, error) {
	switch n.(type) {
	case *column:
		e.upup = n
	case *element:
		e.upup = n
	default:
		return nil, errors.New(e_element_set_vert)
	}
	return n, nil
}

func (e *element) setDn(n node) (node, error) {
	switch n.(type) {
	case *column:
		e.down = n
	case *element:
		e.down = n
	default:
		return nil, errors.New(e_element_set_vert)
	}
	return n, nil
}
