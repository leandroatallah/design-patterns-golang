# Abstract Factory

> The abstract factory pattern provides a way to create families of related objects without imposing their concrete classes, by encapsulating a group of individual factories that have a common theme without specifying their concrete classes. (Wikipedia)

> Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes. (Refactoring Guru)

## Structure

1. **Abstract Products** declare interfaces for a set of distinct but related products which make up a product family.
2. **Concrete Products** are various implementations of abstract products, grouped by variants. Each abstract product must be implemented in all given variants.
3. The **Abstract Factory** interface declares a set of methods for creating each of the abstract products.
4. **Concrete Factories** implement creation methods of the abstract factory. Each concrete factory corresponds to a specific variant of products and creates only those product variants.
5. Although concrete factories instantiate concrete products, signatures of their creation methods must return corresponding abstract products. This way the client code that uses a factory doesn’t get coupled to the specific variant of the product it gets from a factory. The **Client** can work with any concrete factory/product variant, as long as it communicates with their objects via abstract interfaces.

## Example

```go
package main

import "fmt"

// astract factory
type EnemyFactory interface {
 createMinion(name string) IMinion
 createBoss(name string, level int8) IBoss
}

// concrete factory
type LandEnemyFactory struct{}

type WaterEnemyFactory struct{}

func (f *WaterEnemyFactory) createMinion(name string) IMinion {
 return &WaterMinion{name}
}

func (f *WaterEnemyFactory) createBoss(name string, level int8) IBoss {
 return &WaterBoss{name, level}
}

func (f *LandEnemyFactory) createMinion(name string) IMinion {
 return &LandMinion{name}
}

func (f *LandEnemyFactory) createBoss(name string, level int8) IBoss {
 return &LandBoss{name, level}
}

// abstract product
type IMinion interface {
 GetName() string
}
type IBoss interface {
 GetName() string
 GetLevel() int8
}

// concrete product
type LandMinion struct {
 name string
}

func (m *LandMinion) GetName() string {
 return m.name
}

type LandBoss struct {
 name  string
 level int8
}

func (b *LandBoss) GetName() string {
 return b.name
}
func (b *LandBoss) GetLevel() int8 {
 return b.level
}

type WaterMinion struct {
 name string
}

func (m *WaterMinion) GetName() string {
 return m.name
}

type WaterBoss struct {
 name  string
 level int8
}

func (b *WaterBoss) GetName() string {
 return b.name
}
func (b *WaterBoss) GetLevel() int8 {
 return b.level
}

// Creation methods
func NewLandEnemyFactory() EnemyFactory {
 return &LandEnemyFactory{}
}
func NewWaterEnemyFactory() EnemyFactory {
 return &WaterEnemyFactory{}
}

// Print methods
func PrintMinion(minion IMinion) {
 fmt.Printf("Minion\nName: %s\n\n", minion.GetName())
}
func PrintBoss(boss IBoss) {
 fmt.Printf("Boss\nName: %s\nlevel: %d\n\n", boss.GetName(), boss.GetLevel())
}

func main() {
 LandEnemyFactory := NewLandEnemyFactory()
 waterEnemyFactory := NewWaterEnemyFactory()

 smartCat := LandEnemyFactory.createMinion("Smart Cat")
 smartLion := LandEnemyFactory.createBoss("Smart Lion", 2)

 angryFish := waterEnemyFactory.createMinion("Angry Fish")
 angryShark := waterEnemyFactory.createBoss("Angry Shark", 3)

 PrintMinion(smartCat)
 PrintBoss(smartLion)
 PrintMinion(angryFish)
 PrintBoss(angryShark)
}
```
