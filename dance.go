package goDLX

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
	n := new(Matrix)
	n.initRoot()
	return n
}

func (r *Matrix) initRoot() {
	r.L = r
	r.R = r
	r.head = true
}
