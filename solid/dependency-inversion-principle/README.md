# DIP - Dependency Inversion Principle

> Depend upon abstractions. Do not depend upon concrete classes.

- High-level modules should not import anything from low-level modules. Both should depend on abstractions (e.g., interfaces).
- Abstractions should not depend on details. Details (concrete implementations) should depend on abstractions.
- It is difference of Dependency Injection

### ❌ Wrong use

```go
package main

import (
 "fmt"
 "strings"
)

type Skill struct {
 name string
}

type Developer struct {
 name   string
 skills []Skill
}

func (d *Developer) AddSkill(skill *Skill) {
 d.skills = append(d.skills, *skill)
}

type Team struct {
  // Direct depedency on Developer
 employees []Developer
}

func (e *Team) ListEmployees() {
  // Iterates over Developer concrete type
 for _, e := range e.employees {
  fmt.Printf("Name: %s\n", e.name)
  skills := []string{}
  for _, s := range e.skills {
   skills = append(skills, s.name)
  }
  fmt.Printf("Skills: %s\n\n", strings.Join(skills, ", "))
 }
}

func main() {
 dev1 := Developer{"Phill Brooks", []Skill{}}
 dev2 := Developer{"Colt Cabana", []Skill{}}

 skl1 := Skill{"React"}
 skl2 := Skill{"Golang"}

 dev1.AddSkill(&skl1)
 dev1.AddSkill(&skl2)
 dev2.AddSkill(&skl2)

 team1 := Team{[]Developer{dev1, dev2}}
 team1.ListEmployees()
}
```

### ✅ Good use

```go
package main

import (
 "fmt"
 "strings"
)

// Abstraction
type EmployeeLister interface {
 GetName() string
 GetSkills() []string
}

type Skill struct {
 name string
}

type Developer struct {
 name   string
 skills []Skill
}

// Developer implements EmployeeList
func (d *Developer) GetName() string {
 return d.name
}
func (d *Developer) GetSkills() []string {
 result := []string{}
 for _, s := range d.skills {
  result = append(result, s.name)
 }
 return result
}
func (d *Developer) AddSkill(skill *Skill) {
 d.skills = append(d.skills, *skill)
}

type Team struct {
  // Now, Team depends on EmployeeLister abstraction
 employees []EmployeeLister
}

func (e *Team) ListEmployees() {
 // It iterates over abstraction and not concrete type
 for _, e := range e.employees {
  fmt.Printf("Name: %s\n", e.GetName())
  fmt.Printf("Skills: %s\n\n", strings.Join(e.GetSkills(), ", "))
 }
}

func main() {
 dev1 := Developer{"Phill Brooks", []Skill{}}
 dev2 := Developer{"Colt Cabana", []Skill{}}

 skl1 := Skill{"React"}
 skl2 := Skill{"Golang"}

 dev1.AddSkill(&skl1)
 dev1.AddSkill(&skl2)
 dev2.AddSkill(&skl2)

  // It injects abstract types
 team1 := Team{[]EmployeeLister{&dev1, &dev2}}
 team1.ListEmployees()
}
```
