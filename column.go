package goDLX

import "errors"

const (
	ERROR_COLUMN_SET_HORZ = "Only Head or other Column can be left or right of Column"
	ERROR_COLUMN_SET_VERT = "Only Element or Self can be above or below Column"
)

type Column struct {
	lft, rgt, up, dn Node
	name             string
	size             int  // tracks number of 1s in a column
	optional         bool // set for optional columns
}

// NewColumn creates an empty constraint column ready to add to a matrix
func NewColumn(name string, optional bool) *Column {
	c := new(Column)
	c.initColumn(name, optional)
	return c
}

func (c *Column) initColumn(name string, optional bool) {
	c.lft = c
	c.rgt = c
	c.up = c
	c.dn = c
	c.name = name
	c.optional = optional
}

// Lft returns the node to the left of the Column
func (c *Column) Lft() Node {
	return c.lft
}

// SetLft sets a node to the left of the Column node.
func (c *Column) SetLft(n Node) (Node, error) {
	switch n.(type) {
	case *Head:
		// totally allowed
	case *Column:
		if n == c {
			return nil, errors.New(ERROR_COLUMN_SET_HORZ)
		}
	default:
		return nil, errors.New(ERROR_COLUMN_SET_HORZ)
	}
	c.lft = n
	return n, nil
}

// Rgt returns the node to the right of the Column
func (c *Column) Rgt() Node {
	return c.rgt
}

// SetRgt sets the node to the right of the Column
func (c *Column) SetRgt(n Node) (Node, error) {
	switch n.(type) {
	case *Head:
		// totally allowed
	case *Column:
		if n == c {
			return nil, errors.New(ERROR_COLUMN_SET_HORZ)
		}
	default:
		return nil, errors.New(ERROR_COLUMN_SET_HORZ)
	}
	c.rgt = n
	return n, nil
}

// Up returns the node above the Column
func (c *Column) Up() Node {
	return c.up
}

// SetUp returns the Column itself unmodified as Column nodes are only horizontally linked
func (c *Column) SetUp(n Node) (Node, error) {
	switch n.(type) {
	case *Column:
		if n != c {
			return nil, errors.New(ERROR_COLUMN_SET_VERT)
		}
	default:
		return nil, errors.New(ERROR_COLUMN_SET_VERT)
	}
	c.up = n
	return c, nil
}

// Dn returns the Column itself as Column nodes are only horizontally linked
func (c *Column) Dn() Node {
	return c
}

// SetDn returns the Column itself unmodified as Column nodes are only horizontally linked
func (c *Column) SetDn(n Node) (Node, error) {
	switch n.(type) {
	case *Column:
		if n != c {
			return nil, errors.New(ERROR_COLUMN_SET_VERT)
		}
	default:
		return nil, errors.New(ERROR_COLUMN_SET_VERT)
	}
	c.dn = n
	return c, nil
}

// Col returns the Column itself as it is the header in the column list
func (c *Column) Col() Node {
	return c
}
