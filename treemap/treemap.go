package treemap

import (
	"fmt"
)

// node (Private) - Defines the structure for each individual node in a linked TreeMap
type node struct {
	key         string   // key of node
	value       int      // Value of Node
	left, right *node    // Left and right of node
	TreeMap     *TreeMap // Pointer to the TreeMap it is attached to
}

func (n node) compareTo(than string) int {
	var result int
	var small string
	if len(n.key) < len(than) {
		small = n.key
	} else {
		small = than
	}
	for x := 0; x < len(small); x++ {
		result = int(n.key[x]) - int(than[x])
		if result == 0 {
			continue
		} else {
			break
		}
	}
	return result
}

func (n node) String() string {
	return fmt.Sprintf("%v=%v", n.key, n.value)
}

// TreeMap (Public) - The container for all the linked nodes in a set
type TreeMap struct {
	root *node // the begining node
	size int   // size of the TreeMap
}

// New (Public) - Returns an initialized TreeMap.
func New() *TreeMap { return new(TreeMap) }

// Add (Public) - if key in map, replaces old value with new and returns old value
// otherwise adds key value pair and returns nil
func (t *TreeMap) Add(key string, value int) int {
	if key == "" || value == 0 {
		return -1
	}
	new := &node{key: key, value: value}
	if t.root == nil {
		t.root = new
		t.size++
		return t.root.value
	}
	current := t.root
	var parent *node
	for current != nil {
		result := new.compareTo(current.key)
		if result == 0 {
			return -1
		}
		parent = current
		if result < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	if new.compareTo(parent.key) < 0 {
		parent.left = new
	} else {
		parent.right = new
	}
	t.size++
	return new.value
}

// Remove (Public) - removes key value pair from map and returns value
// if key not found, returns nil
func (t *TreeMap) Remove(key string) int {
	/* 	if key == nil {
	   		return nil
	   	}
	   	current := &l.head
	   	for i := 0; i < l.size; i++ {
	   		if current.next.key == key {
	   			removed := current.next.value
	   			current.next = current.next.next
	   			l.size--
	   			return removed
	   		}
	   		current = current.next
	   	} */
	return -1
}

// Get (Public) - returns value associated with key
// if key not found, returns null
func (t *TreeMap) Get(key string) int {
	current := t.root
	for {
		if current == nil {
			return -1
		}
		//fmt.Println(current.key)
		result := current.compareTo(key)
		if result == 0 {
			break
		} else if result > 0 {
			//fmt.Println("Moving left...")
			current = current.left
		} else {
			//fmt.Println("Moving right...")
			current = current.right
		}
	}
	return current.value
}

// Contains (Public) - returns true if key is in map
func (t *TreeMap) Contains(key string) bool {
	current := t.root
	for {
		if current == nil {
			return false
		}
		//fmt.Println(current.key)
		result := current.compareTo(key)
		if result == 0 {
			break
		} else if result > 0 {
			//fmt.Println("Moving left...")
			current = current.left
		} else {
			//fmt.Println("Moving right...")
			current = current.right
		}
	}
	return true
}

// Size (Public) - Returns the size of the TreeMap
func (t *TreeMap) Size() int {
	return t.size
}

// String (Public) - Allows for the fmt.Print* functions to print the TreeMap struct as a string.
func (t *TreeMap) String() string {
	if t.size == 0 {
		return "{}"
	}
	b := ""
	s := "{" + inorder(t.root, &b)
	return s + "}"
}

func inorder(n *node, s *string) string {
	if n.left != nil {
		inorder(n.left, s)
	}
	*s += n.String() + " "
	if n.right != nil {
		inorder(n.right, s)
	}
	return *s
}
