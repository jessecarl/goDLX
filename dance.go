package goDLX

import (
	"errors"
)

// Matrix is the fundamental unit in the Algorithm X implementation
// as described in [Knuth, Donald (2000). "Dancing Links". _Millenial Perspectives in Computer Science_. P159 *187*.
// Name and size only apply to columns
type Matrix struct {
	L, R, U, D, C *Matrix
	Name          string
	size          int  // count of 1s in the column
	optional      bool // optional columns do not have to be satisfied but can be only once
	head          bool // checked for the root or head node
}

// New returns an empty matrix. This creates a single head or root node to which
// all other nodes are linked. According to the algorithm, only the left and right
// values are used for the root element.
func New() *Matrix {
	r := new(Matrix)
	r.initRoot()
	return r
}

func (r *Matrix) initRoot() {
	r.L = r
	r.R = r
	r.head = true
}

func (r *Matrix) isRoot() bool {
	return r.head
}

// AddCol adds an empty constraint column to the matrix, returning the column pointer.
// Columns are added to the left of the head node. This function can only be called on
// the root node.
func (r *Matrix) AddCol(name string, optional bool) (*Matrix, error) {
	if !r.isRoot() {
		return nil, errors.New("Not a root element")
	}
	c := new(Matrix)
	c.initCol(name, optional)
	c.L = r.L
	c.R = r
	r.L = c
	c.L.R = c
	return c, nil
}

func (c *Matrix) initCol(name string, optional bool) {
	c.U = c
	c.D = c
	c.C = c
	c.Name = name
	c.size = 0
	c.optional = optional
}
