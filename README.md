Package
===

package goDLX

    import "github.com/jessecarl/goDLX"

Package goDLX provides a Go implementation of Dancing Links Algorithm X
by Donald Knuth.

I'm attempting to follow the paper as exactly as possible to start. I
can cater it a bit more to to [Go](http://golang.org) way of solving the
problem once the initial implementation is complete.

Types
===

    type Head struct {
        // contains filtered or unexported fields
    }

Head nodes are the master column headers. These nodes form the heart of
each sparse matrix

    func New() *Head

New starts a new sparse matrix by creating the Head node.

    func (h *Head) AddCol(name string, optional bool) error

AddCol adds a Column to the matrix. Each column has a name for reference
and can be set to optional. Where required columns have one and only one
valid row, optional columns have zero or one valid row.

    func (h *Head) AddRow(cols []string) error

AddRow adds a row to the matrix Each row is populated by adding links
for every column named in the cols array. Once rows are added, the
matrix is locked and no further columns may be added.


