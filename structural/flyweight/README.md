# Flyweight

> Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object. (Refactoring Guru)

> Flyweight design pattern refers to an object that minimizes memory usage by sharing some of its data with other similar objects. (Wikipedia)

## Structure

1. The Flyweight pattern is merely an optimization. Before applying it, make sure your program does have the RAM consumption problem related to having a massive number of similar objects in memory at the same time. Make sure that this problem can’t be solved in any other meaningful way.
2. The **Flyweight** class contains the portion of the original object’s state that can be shared between multiple objects. The same flyweight object can be used in many different contexts. The state stored inside a flyweight is called intrinsic. The state passed to the flyweight’s methods is called extrinsic.

## Examples

### Basic example

```go
// flyweight.go
package main

import "fmt"

// Dados compartilhados (flyweight)
type EnemyType struct {
 Name   string
 Sprite string // caminho do sprite
 Speed  float64
}

// Dados únicos de cada inimigo
type Enemy struct {
 Type *EnemyType // referência ao flyweight
 X, Y float64
 HP   int
}

func main() {
 // Flyweights
 zombieType := &EnemyType{Name: "Zombie", Sprite: "zombie.png", Speed: 1.0}
 batType := &EnemyType{Name: "Bat", Sprite: "bat.png", Speed: 2.5}

 enemies := []*Enemy{
  {Type: zombieType, X: 10, Y: 20, HP: 100},
  {Type: zombieType, X: 15, Y: 25, HP: 100},
  {Type: batType, X: 5, Y: 8, HP: 50},
 }

 for _, e := range enemies {
  fmt.Printf("%s at (%.0f,%.0f) HP:%d\n", e.Type.Name, e.X, e.Y, e.HP)
 }
}

```

### Text formatting

```go
package main

import (
 "fmt"
 "strings"
 "unicode"
)

// It is a anti-flyweight example
type FormattedText struct {
 plainText  string
 capitalize []bool // same size of plainText
}

func (f *FormattedText) String() string {
 sb := strings.Builder{}
 for i := 0; i < len(f.plainText); i++ {
  c := f.plainText[i]
  if f.capitalize[i] {
   sb.WriteRune(unicode.ToUpper(rune(c)))
  } else {
   sb.WriteRune(rune(c))
  }
 }
 return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
 for i := start; i <= end; i++ {
  f.capitalize[i] = true
 }
}

func NewFormattedText(plainText string) *FormattedText {
 return &FormattedText{plainText: plainText, capitalize: make([]bool, len(plainText))}
}

// It is a Flyweight
type TextRange struct {
 Start, End               int
 Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
 return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
 plainText  string
 formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
 return &BetterFormattedText{plainText: plainText}
}

// Fluent function
func (b *BetterFormattedText) Range(start, end int) *TextRange {
 r := &TextRange{start, end, false, false, false}
 b.formatting = append(b.formatting, r)
 return r
}

func (b *BetterFormattedText) String() string {
 sb := strings.Builder{}
 for i := 0; i < len(b.plainText); i++ {
  c := b.plainText[i]
  for _, r := range b.formatting {
   if r.Covers(i) && r.Capitalize {
    c = uint8(unicode.ToUpper(rune(c)))
   }
  }
  sb.WriteRune(rune(c))
 }
 return sb.String()
}

func main() {
 text := "The new world order"

 // Bad implementation
 ft := NewFormattedText(text)
 ft.Capitalize(8, 12)
 fmt.Println(ft.String())

 // FLyweight implementation
 bft := NewBetterFormattedText(text)
 bft.Range(8, 12).Capitalize = true
 fmt.Println(bft.String())
}
```

### User names

```go
package main

import (
 "fmt"
 "strings"
)

// Without memory saving
type User struct {
 fullName string
}

func (u *User) FullName() string {
 return u.fullName
}

func NewUser(fullName string) *User {
 return &User{fullName}
}

// With memory saving
var allNames []string

type User2 struct {
 names []uint8
}

func NewUser2(fullName string) *User2 {
 getOrAdd := func(s string) uint8 {
  for i := range allNames {
   if allNames[i] == s {
    return uint8(i)
   }
  }
  allNames = append(allNames, s)
  return uint8(len(allNames) - 1)
 }

 result := User2{}
 parts := strings.Split(fullName, " ")
 for _, p := range parts {
  result.names = append(result.names, getOrAdd(p))
 }
 return &result
}

func (u *User2) FullName() string {
 var parts []string
 for _, id := range u.names {
  parts = append(parts, allNames[id])
 }
 return strings.Join(parts, " ")
}

func main() {
 john := NewUser("John Doe")
 fmt.Println(john.FullName())

 john2 := NewUser2("John Doe")
 fmt.Println(john2, john2.FullName())
}
```
