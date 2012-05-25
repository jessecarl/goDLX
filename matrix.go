package goDLX

// The node interface navigates a sparse matrix.
// Any type implementing a 2-dimensional doubly-linked list satisfies this.
// (despite the convention being that interfaces have the -er suffix, this special case
// can't do that without being entirely unclear)
type node interface {
	lft() node
	setLft(node) (node, error)
	rgt() node
	setRgt(node) (node, error)
	up() node
	setUp(node) (node, error)
	dn() node
	setDn(node) (node, error)
	col() node
}

// linkHorz links a *new* left node to an *existing* right node as shown:
//     +---+  +---+  +---+
//     |   |  |   |  |   |
//     |   |  | a |  | b |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |   |  | a |  | b |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  | a |  | b |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  | a |  | b |
//     |   |<-|   |<-|   |
//     +---+  +---+  +---+
// Only need a single horizontal link adding method because these are circular links.
// If one wants to add a node to the right instead of left, simply use
// linkHorz( a , rgt.rgt())
func linkHorz(a, b node) error {
	if a == b {
		// not an error, but nothing happens
		return nil
	}
	if _, err := a.setLft(b.lft()); err != nil {
		return err
	}
	if _, err := a.setRgt(b); err != nil {
		return err
	}
	if _, err := a.lft().setRgt(a); err != nil {
		return err
	}
	if _, err := a.rgt().setLft(a); err != nil {
		// undo our modification to the existing list
		a.lft().setRgt(b)
		return err
	}
	return nil
}

// linkVert links a *new* up node to an *existing* down node as shown:
//     +---+  +---+  +---+
//     |   |  |   |  |   |
//     |   |  | a |  | b |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |   |  | a |  | b |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  | a |  | b |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  | a |  | b |
//     |   |<-|   |<-|   |
//     +---+  +---+  +---+
// This is the vertical analog to the linkHorz() function.
func linkVert(a, b node) error {
	if a == b {
		// not an error, but nothing happens
		return nil
	}
	if _, err := a.setUp(b.up()); err != nil {
		return err
	}
	if _, err := a.setDn(b); err != nil {
		return err
	}
	if _, err := a.up().setDn(a); err != nil {
		return err
	}
	if _, err := a.dn().setUp(a); err != nil {
		// undo our modification to the existing list
		a.up().setDn(b)
		return err
	}
	return nil
}
