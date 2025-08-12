# Iterator

> Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.). (Refactoring Guru)

> The iterator pattern is a design pattern in which an iterator is used to traverse a container and access the container's elements. The iterator pattern decouples algorithms from containers; in some cases, algorithms are necessarily container-specific and thus cannot be decoupled. (Wikipedia)

## Examples

### Basic example

```go
package main

import "fmt"

// Model
type Person interface {
 Name() string
}
type Employee struct {
 name string
}

func (e *Employee) Name() string {
 return e.name
}

// iterator interface
type Iterator interface {
 GetNext() (Person, bool)
 HasMore() bool
}

// Concrete iterator
type EmployeeIterator struct {
 collection      *EmployeeCollection
 currentPosition int
}

func NewPersonIterator(collection *EmployeeCollection) *EmployeeIterator {
 return &EmployeeIterator{collection: collection}
}

func (i *EmployeeIterator) GetNext() (Person, bool) { // change return type
 if i.HasMore() {
  result := i.collection.employeeList[i.currentPosition]
  i.currentPosition++
  return result, true
 }
 return nil, false
}
func (i *EmployeeIterator) HasMore() bool {
 return i.currentPosition < len(i.collection.employeeList)
}

// iterable collection
type IterableCollection interface {
 CreateIterator() Iterator
}

// concrete collection
type EmployeeCollection struct {
 employeeList []Person
}

func NewEmployeeCollection(employees []Person) *EmployeeCollection {
 return &EmployeeCollection{employeeList: employees}
}

func (c *EmployeeCollection) CreateIterator() Iterator {
 return NewPersonIterator(c)
}

func main() {
 employees := []Person{
  &Employee{"Leonard"},
  &Employee{"Raphael"},
 }
 iterator := NewEmployeeCollection(employees).CreateIterator()

 for iterator.HasMore() {
  if next, exists := iterator.GetNext(); exists {
   fmt.Println(next.Name())
  }
 }
}
```

### BInary tree iterator

```go
package main

import "fmt"

type Node struct {
 Value               int
 left, right, parent *Node
}

func NewNode(value int, left, right *Node) *Node {
 n := &Node{Value: value, left: left, right: right}
 left.parent = n
 right.parent = n
 return n
}

func NewTerminalNode(value int) *Node {
 return &Node{Value: value}
}

type InOrderIterator struct {
 Current       *Node
 root          *Node
 returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
 i := &InOrderIterator{root, root, false}
 for i.Current.left != nil {
  i.Current = i.Current.left
 }
 return i
}

func (i *InOrderIterator) Reset() {
 i.Current = i.root
 i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
 if i.Current == nil {
  return false
 }
 if !i.returnedStart {
  i.returnedStart = true
  return true
 }

 if i.Current.right != nil {
  i.Current = i.Current.right
  for i.Current.left != nil {
   i.Current = i.Current.left
  }
  return true
 } else {
  p := i.Current.parent
  for p != nil && i.Current == p.right {
   i.Current = p
   p = p.parent
  }
  i.Current = p
  return i.Current != nil
 }
}

type BinaryTree struct {
 root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
 return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
 return NewInOrderIterator(b.root)
}

func main() {
 root := NewNode(1,
  NewTerminalNode(2),
  NewTerminalNode(3),
 )

 t := NewBinaryTree(root)
 for i := t.InOrder(); i.MoveNext(); {
  fmt.Printf("%d\n", i.Current.Value)
 }
}
```
