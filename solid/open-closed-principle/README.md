# OCP - Open-Closed Principle

> Open for extension, closed for modification.

- The goal of OCP is to allow classes to be easily extended to incorporate new behavior without modifying existing code.
- While it may seem like a contradiction, there are techniques for allowing code to be extended without direct modification, as Enterprise pattern Specification.

<aside>
⚠️

Applying the OCP EVERYWHERE is wasteful, unnecessary, and can lead to complex, hard to understand code.

</aside>

### ❌ Wrong use

```go
package main

import "fmt"

type Type int

const (
 Car Type = iota
 Motorcycle
 Bicycle
)

type Vehicle struct {
 Type  Type
 Price float64
}

// If you need to add a new type of vehicle,
// you need to modify the original code and add a new condition.
func (v *Vehicle) GetTotalWithTax() float64 {
 if v.Type == Car {
  return v.Price + (v.Price * 0.2)
 }
 if v.Type == Motorcycle {
  return v.Price + (v.Price * 0.15)
 }
 if v.Type == Bicycle {
  return v.Price + (v.Price * 0.1)
 }
 return v.Price
}

func main() {
 car := Vehicle{Car, 15_000}
 moto := Vehicle{Motorcycle, 12_000}
 bicycle := Vehicle{Bicycle, 2_000}

 carTotal := car.GetTotalWithTax()
 motoTotal := moto.GetTotalWithTax()
 bicycleTotal := bicycle.GetTotalWithTax()

 fmt.Println("Car Total Price: ", carTotal)
 fmt.Println("Motorcycle Total Price: ", motoTotal)
 fmt.Println("Bicycle Total Price: ", bicycleTotal)
}
```

### ✅ Good use

```go
package main

import "fmt"

type Type int

const (
 Car Type = iota
 Motorcycle
 Bicycle
)

type Vehicle struct {
 Type  Type
 Price float64
}

type TotalTaxCounter interface {
 GetTotalWithTax(t Type, p float64) float64
}

// You can add new vehicle condition,
// without modify the original Vehicle code.
type CarPricer struct{}
func (c CarPricer) GetTotalWithTax(v *Vehicle) float64 {
 return v.Price + (v.Price * 0.2)
}

type MotoPricer struct{}
func (c MotoPricer) GetTotalWithTax(v *Vehicle) float64 {
 return v.Price + (v.Price * 0.15)
}
type BicyclePricer struct{}

func (c BicyclePricer) GetTotalWithTax(v *Vehicle) float64 {
 return v.Price + (v.Price * 0.1)
}

func main() {
 car := Vehicle{Car, 15_000}
 moto := Vehicle{Motorcycle, 12_000}
 bicycle := Vehicle{Bicycle, 2_000}

 carPricer := CarPricer{}
 carTotal := carPricer.GetTotalWithTax(&car)

 motoPricer := MotoPricer{}
 motoTotal := motoPricer.GetTotalWithTax(&moto)

 bicyclePricer := BicyclePricer{}
 bicycleTotal := bicyclePricer.GetTotalWithTax(&bicycle)

 fmt.Println("Car Total Price: ", carTotal)
 fmt.Println("Motorcycle Total Price: ", motoTotal)
 fmt.Println("Bicycle Total Price: ", bicycleTotal)
}
```
