# Prototype

> Prototype is used when the types of objects to create is determined by a prototypical instance, which is cloned to produce new objects (Wikipedia)

> Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes. (Refactoring Guru)

## Structure

1. The **Prototype** interface declares the cloning methods. In most cases, it’s a single clone method.
2. The **Concrete Prototype** class implements the cloning method. In addition to copying the original object’s data to the clone, this method may also handle some edge cases of the cloning process related to cloning linked objects, untangling recursive dependencies, etc.
3. The **Client** can produce a copy of any object that follows the prototype interface.

## Deep Copying

A deep copy of an object creates a new instance with properties that don't share references with the source object. This ensures changes to either the source or copy won't affect the other. This differs from shallow copies, where modifying nested properties in one object can unintentionally change the other.

### Deep copy

```go
package main

import (
 "bytes"
 "encoding/gob"
 "fmt"
)

type Address struct {
 StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
  return &Address{
   a.StreetAddress,
   a.City,
   a.Country,
  }
}

type Person struct {
 Name    string
 Address *Address
 Friends []string
}

func (p *Person) DeepCopy() *Person {
 q := *p
 q.Address = p.Address.DeepCopy()
  copy(q.Friends, p.Friends)
 return &q
}

func main() {
 bob := Person{"Bob", &Address{"123 Street", "Fundão", "Portugal"}, []string{"Luke", "Matthew"}}
 clark := bob.DeepCopy()
 clark.Name = "Clark"
 clark.Address.StreetAddress = "321 Another Street"
 clark.Friends = append(clark.Friends, "Holy")

 fmt.Println(bob, bob.Address)
 fmt.Println(clark, clark.Address)
}

```

### Deep Copy by serialization using encoding/gob

```go
package main

import (
 "bytes"
 "encoding/gob"
 "fmt"
)

type Address struct {
 StreetAddress, City, Country string
}

type Person struct {
 Name    string
 Address *Address
 Friends []string
}

func (p *Person) DeepCopy() *Person {
 b := bytes.Buffer{}
 e := gob.NewEncoder(&b)
 _ = e.Encode(p)
 d := gob.NewDecoder(&b)
 result := Person{}
 _ = d.Decode(&result)
 return &result
}

func main() {
 bob := Person{"Bob", &Address{"123 Street", "Fundão", "Portugal"}, []string{"Luke", "Matthew"}}
 clark := bob.DeepCopy()
 clark.Name = "Clark"
 clark.Address.StreetAddress = "321 Another Street"
 clark.Friends = append(clark.Friends, "Holy")

 fmt.Println(bob, bob.Address)
 fmt.Println(clark, clark.Address)
}
```

### Prototype Factory example

A prototype factory provides a convenient API for using prototypes

```go
package main

import (
 "bytes"
 "encoding/gob"
 "fmt"
)

type Address struct {
 Suite               int
 StreetAddress, City string
}

type Employee struct {
 Name   string
 Office Address
}

func (p *Employee) DeepCopy() *Employee {
 // no error handling below
 b := bytes.Buffer{}
 e := gob.NewEncoder(&b)
 _ = e.Encode(p)
 d := gob.NewDecoder(&b)
 result := Employee{}
 d.Decode(&result)
 return &result
}

var mainOffice = Employee{"", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee{"", Address{0, "66 West Dr", "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee {
 result := proto.DeepCopy()
 result.Name = name
 result.Office.Suite = suite
 return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
 return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
 return newEmployee(&auxOffice, name, suite)
}

func main() {
 john := NewMainOfficeEmployee("John", 100)
 jane := NewAuxOfficeEmployee("Jane", 200)

 fmt.Println(john)
 fmt.Println(jane)
}
```
