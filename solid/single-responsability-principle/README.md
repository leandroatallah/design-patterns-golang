# Single Responsibility Principle

- A class should have only one reason to change.
- Classes that adhere to this principle tend to have high cohesion. High cohesion means that a class is designed around a set of related functions and have a single purpose.
- Could be called as **Separation of concerns**.
- The “one thing” in this principle could be misunderstood. It refers to a “single reason to change” and not a single method or a single operation. The single responsibility follows a stakeholder perspective. An update to a class, should not affect other areas.
- Anti-pattern God object:
  _In contrast to SRP, most of such a program's overall functionality is coded into a single "all-knowing" object, which maintains most of the information about the entire program, and also provides most of the methods for manipulating this data. Because this object holds so much data and requires so many methods, its role in the program becomes god-like (all-knowing and all-encompassing)._

### ❌ Wrong use

```go
type Stock struct {
 products []string
}

func (s *Stock) AddProduct(text string) {
 s.products = append(s.products, text)
}

func (s *Stock) RemoveProduct(text string) {
 // ...
}

func (s *Stock) String() string {
 return strings.Join(s.products, "\n")
}

// ❌ This method breaks the SRP
func (s *Stock) save(filename string) {
 _ = ioutil.WriteFile(filename, []byte(
}
// ❌ This method breaks the SRP
func (s *Stock) Load(filename string) {
 // ...
}
// ❌ This method breaks the SRP
func (s *Stock) LoadFromWeb(furl *url.URL) {
 // ...
}
```

### ✅ Good use: Method 1

```go
type Stock struct {
 products []string
}

func (s *Stock) AddProduct(text string) {
 s.products = append(s.products, text)
}

func (s *Stock) RemoveProduct(text string) {
 // ...
}

func (s *Stock) String() string {
 return strings.Join(s.products, "\n")
}

// ✅ Using a method that exists by itself, out of Stock
var LineSeparator = "\n"
func SaveToFile(s *Stock, filename string) {
 _ = ioutil.WriteFile(filename, []byte(
  strings.Join(s.products, LineSeparator)
 ), 0644)
}
```

### ✅ Good use: Method 2

```go
type Stock struct {
 products []string
}

func (s *Stock) AddProduct(text string) {
 s.products = append(s.products, text)
}

func (s *Stock) RemoveProduct(text string) {
 // ...
}

func (s *Stock) String() string {
 return strings.Join(s.products, "\n")
}

// ✅ Using a separated module for persistence
type StockPersistence struct {
 lineSeparator string
}
func (p *StockPersistence) Save(filename string) {
 _ = ioutil.WriteFile(filename, []byte(
  strings.Join(s.products, p.lineSeparator)
 ), 0644)
}
```
