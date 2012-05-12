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

// LinkLft links a new left Node to an existing right Node as shown:
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
func LinkLft(lft, rgt Node) error {
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

// LinkRgt links a new right Node to an existing left Node as shown:
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |lft|  |rgt|  |   |
//     |   |  |   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |lft|  |rgt|  |   |
//     |   |<-|   |  |   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |  |   |->|   |
//     |lft|  |rgt|  |   |
//     |   |<-|   |<-|   |
//     +---+  +---+  +---+
//     +---+  +---+  +---+
//     |   |->|   |->|   |
//     |lft|  |rgt|  |   |
//     |   |<-|   |<-|   |
//     +---+  +---+  +---+
func LinkRgt(rgt, lft Node) error {
	if _, err := rgt.SetRgt(lft.Rgt()); err != nil {
		return err
	}
	if _, err := rgt.SetLft(lft); err != nil {
		return err
	}
	if _, err := rgt.Rgt().SetLft(rgt); err != nil {
		return err
	}
	if _, err := rgt.Lft().SetRgt(rgt); err != nil {
		// undo our modification to the existing list
		rgt.Rgt().SetLft(lft)
		return err
	}
	return nil
}
