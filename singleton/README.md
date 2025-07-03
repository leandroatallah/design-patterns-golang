# Singleton

> Singleton pattern restricts the instantiation of a class to a singular instance. The pattern is useful when exactly one object is needed to coordinate actions across a system. (Wikipedia)

> Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance. (Refactoring Guru)

## Concept

1. **Ensure that a class has just a single instance**: imagine that you created an object, but after a while decided to create a new one. Instead of receiving a fresh object, you’ll get the one you already created
2. **Provide a global access point to that instance**: Just like a global variable, the Singleton pattern lets you access some object from anywhere in the program. However, it also protects that instance from being overwritten by other code.

## Examples

### Mutex example

```go
package main

import (
 "fmt"
 "sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() *single {
 if singleInstance == nil {
  lock.Lock()
  defer lock.Unlock()
  if singleInstance == nil {
   fmt.Println("Creating single instance now.")
   singleInstance = &single{}
  } else {
   fmt.Println("Single instance already created.")
  }
 } else {
  fmt.Println("Single instance already created. 1")
 }

 return singleInstance
}

func main() {
 for i := 0; i < 30; i++ {
  go GetInstance()
 }

 // Prevent the application to finish.
 fmt.Scanln()
}

```

### Sync Once example

```go
package main

import (
 "fmt"
 "sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
 if singleInstance == nil {
  once.Do(func() {
   fmt.Println("Creating single instance now.")
   singleInstance = &single{}
  })
 } else {
  fmt.Println("Single instance already created.")
 }

 return singleInstance
}

func main() {
 for i := 0; i < 30; i++ {
  go getInstance()
 }

 fmt.Scanln()
}

```
