package goDLX

// The Node interface navigates a sparse matrix.
// Any type implementing a 2-dimensional doubly-linked list satisfies this.
// (despite the convention being that interfaces have the -er suffix, this special case
// can't do that without being entirely unclear)
type Node interface {
	Lft() Node
	SetLft(Node) (Node, error)
	Rgt() Node
	SetRgt(Node) (Node, error)
	Up() Node
	SetUp(Node) (Node, error)
	Dn() Node
	SetDn(Node) (Node, error)
	Col() Node
}

// LinkHorz links a *new* left Node to an *existing* right Node as shown:
//     +---+  +---+  +---+
//     |   |  |   |  |   |
//     |   |  |lft|  |rgt|
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |   |  |lft|  |rgt|
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  |lft|  |rgt|
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  |lft|  |rgt|
//     |   |<-|   |<-|   |
//     +---+  +---+  +---+
// Only need a single horizontal link adding method because these are circular links.
// If one wants to add a Node to the right instead of left, simply use
// LinkHorz(lft, rgt.Rgt())
func LinkHorz(lft, rgt Node) error {
	if lft == rgt {
		// not an error, but nothing happens
		return nil
	}
	if _, err := lft.SetLft(rgt.Lft()); err != nil {
		return err
	}
	if _, err := lft.SetRgt(rgt); err != nil {
		return err
	}
	if _, err := lft.Lft().SetRgt(lft); err != nil {
		return err
	}
	if _, err := lft.Rgt().SetLft(lft); err != nil {
		// undo our modification to the existing list
		lft.Lft().SetRgt(rgt)
		return err
	}
	return nil
}

// LinkVert links a *new* up Node to and *existing* down Node as shown:
//     +---+  +---+  +---+
//     |   |  |   |  |   |
//     |   |  |up |  |dn |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |   |  |up |  |dn |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  |up |  |dn |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |   |  |up |  |dn |
//     |   |<-|   |<-|   |
//     +---+  +---+  +---+
// This is the vertical analog to the LinkHorz() function.
func LinkVert(up, dn Node) error {
	if up == dn {
		// not an error, but nothing happens
		return nil
	}
	if _, err := up.SetUp(dn.Up()); err != nil {
		return err
	}
	if _, err := up.SetDn(dn); err != nil {
		return err
	}
	if _, err := up.Up().SetDn(up); err != nil {
		return err
	}
	if _, err := up.Dn().SetUp(up); err != nil {
		// undo our modification to the existing list
		up.Up().SetDn(dn)
		return err
	}
	return nil
}
