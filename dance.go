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
