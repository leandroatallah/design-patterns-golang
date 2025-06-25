# ISP - Interface Segregation Principle

> Clients should not be forced to implement methods it doesn't need

- Specific interfaces prevent classes from implementing unnecessary methods.
- Following the first three principles of SOLID helps to not violate ISP.

### ❌ Wrong use

```go
package main

import "fmt"

type IMachine interface {
 Print()
 Scan()
}

type ModernPrinter struct{}

func (m *ModernPrinter) Print() {
 fmt.Println("Printing...")
}

func (m *ModernPrinter) Scan() {
 fmt.Println("Scanning...")
}

type OldPrinter struct{}

func (o *OldPrinter) Print() {
 fmt.Println("Printing...")
}

// It violates ISP.
func (o *OldPrinter) Scan() {
 panic("This machine does not support this action.")
}

func main() {
 m := ModernPrinter{}
 m.Print()
 m.Scan()

 o := OldPrinter{}
 o.Print()
 o.Scan()
}
```

### ✅ Good use

```go
package main

import "fmt"

type Printer interface {
 Print()
}

type Scanner interface {
 Scan()
}

type MultiFunctionDevice interface {
 Printer
 Scanner
}

type ModernPrinter struct{}

func (m *ModernPrinter) Print() {
 fmt.Println("Printing...")
}

func (m *ModernPrinter) Scan() {
 fmt.Println("Scanning...")
}

type OldPrinter struct{}

func (o *OldPrinter) Print() {
 fmt.Println("Printing...")
}

func main() {
 m := ModernPrinter{}
 m.Print()
 m.Scan()

 o := OldPrinter{}
 o.Print()
}
```
