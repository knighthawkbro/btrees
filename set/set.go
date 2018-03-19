package set

import "fmt"

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	data        string // Value of Node
	right, left *node  // Pointers to the next left or right node
}

func (n node) compareTo(than string) int {
	// Corner case to shortcut if two strings are equal
	if n.data == than {
		return 0
	}
	var result int
	var small string
	if len(n.data) < len(than) {
		small = n.data
	} else {
		small = than
	}
	for x := 0; x < len(small); x++ {
		result = int(n.data[x]) - int(than[x])
		if result == 0 {
			continue
		} else {
			break
		}
	}
	// Had to fix a corner case where the smaller word was just
	// the first part of the Than word.
	if result == 0 {
		if small == n.data {
			return -1
		}
		return 1
	}
	return result
}

// String (Public) -
func (n node) String() string {
	return fmt.Sprint(n.data)
}

// TreeSet (Public) - The container for all the linked nodes in a set
type TreeSet struct {
	root *node // the begining node
	size int   // size of the list
}

// New (Public) - Returns an initialized TreeMap.
func New() *TreeSet { return new(TreeSet) }

// Size (Public) - Returns the length variable for the list as an integer
func (t *TreeSet) Size() int { return t.size }

// Add (Public) - Returns the node in a singly linked list, just adds to the front of the list
func (t *TreeSet) Add(item string) bool {
	if item == "" {
		return false
	}
	new := &node{data: item}
	if t.root == nil {
		t.root = new
		t.size++
		return true
	}
	current := t.root
	var parent *node
	for current != nil {
		result := new.compareTo(current.data)
		if result == 0 {
			return false
		}
		parent = current
		if result < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	if new.compareTo(parent.data) < 0 {
		parent.left = new
	} else {
		parent.right = new
	}
	t.size++
	return true
}

// Remove (Public) - Removes the first item on a list and returns it
func (t *TreeSet) Remove() string {
	return "<nil>"
}

// Get (Public) - Returns the first item list
func (t *TreeSet) Get() string {
	if t.size == 0 {
		return "<nil>"
	}
	return ""
}

// Contains (Public) - Returns true or false whether an item was contained in the list
func (t *TreeSet) Contains(item string) bool {
	current := t.root
	for current != nil {
		result := current.compareTo(item)
		if result == 0 {
			return true
		}
		if result < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

// String (Public) - Allows for the fmt.Print* functions to print the TreeMap struct as a string.
func (t *TreeSet) String() string {
	if t.root == nil {
		return "{}"
	}
	s := "{ "
	inorder(t.root, &s)
	return s + "}"
}

func inorder(n *node, s *string) {
	if n.left != nil {
		inorder(n.left, s)
	}
	*s += n.String() + " "
	if n.right != nil {
		inorder(n.right, s)
	}
}
