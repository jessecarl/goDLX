package goDLX

import "errors"

const (
	e_column_set_horz = "Only Head or other column can be left or right of column"
	e_column_set_vert = "Only Element or Self can be above or below column"
)

type column struct {
	left, right, upup, down node
	name                    string
	size                    int  // tracks number of 1s in a column
	optional                bool // set for optional columns
}

// newColumn creates an empty constraint column ready to add to a matrix
func newColumn(name string, optional bool) *column {
	c := new(column)
	c.initColumn(name, optional)
	return c
}

func (c *column) initColumn(name string, optional bool) {
	c.left = c
	c.right = c
	c.upup = c
	c.down = c
	c.name = name
	c.optional = optional
}

// lft returns the node to the left of the column
func (c *column) lft() node {
	return c.left
}

// setLft sets a node to the left of the column node.
func (c *column) setLft(n node) (node, error) {
	switch n.(type) {
	case *Head:
		// totally allowed
	case *column:
		if n == c {
			return nil, errors.New(e_column_set_horz)
		}
	default:
		return nil, errors.New(e_column_set_horz)
	}
	c.left = n
	return n, nil
}

// rgt returns the node to the right of the column
func (c *column) rgt() node {
	return c.right
}

// setRgt sets the node to the right of the column
func (c *column) setRgt(n node) (node, error) {
	switch n.(type) {
	case *Head:
		// totally allowed
	case *column:
		if n == c {
			return nil, errors.New(e_column_set_horz)
		}
	default:
		return nil, errors.New(e_column_set_horz)
	}
	c.right = n
	return n, nil
}

// up returns the node above the column
func (c *column) up() node {
	return c.upup
}

// setUp returns the column itself unmodified as column nodes are only horizontally linked
func (c *column) setUp(n node) (node, error) {
	switch n.(type) {
	case *column:
		if n != c {
			return nil, errors.New(e_column_set_vert)
		}
	case *element:
		// totally allowed
	default:
		return nil, errors.New(e_column_set_vert)
	}
	c.upup = n
	return n, nil
}

// dn returns the column itself as column nodes are only horizontally linked
func (c *column) dn() node {
	return c.down
}

// setDn returns the column itself unmodified as column nodes are only horizontally linked
func (c *column) setDn(n node) (node, error) {
	switch n.(type) {
	case *column:
		if n != c {
			return nil, errors.New(e_column_set_vert)
		}
	case *element:
		// totally allowed
	default:
		return nil, errors.New(e_column_set_vert)
	}
	c.down = n
	return n, nil
}

// col returns the column itself as it is the header in the column list
func (c *column) col() node {
	return c
}

func (c *column) count() int {
	return c.size
}

func (c *column) label() string {
	return c.name
}

func (c *column) newElement() *element {
	e := new(element)
	e.init(c)
	c.size++
	return e
}
