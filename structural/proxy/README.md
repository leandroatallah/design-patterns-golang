# Proxy

> Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object. (Refactoring Guru)

> A proxy, in its most general form, is a class functioning as an interface to something else. The proxy could interface to anything: a network connection, a large object in memory, a file, or some other resource that is expensive or impossible to duplicate. (Wikipedia)

## Explanation

- The Proxy pattern suggests that you create a new proxy class with the same interface as an original service object.
- Then you update your app so that it passes the proxy object to all of the original object’s clients.
- Upon receiving a request from a client, the proxy creates a real service object and delegates all the work to it.

## Structure

1. The Service Interface declares the interface of the Service. The proxy must follow this interface to be able to disguise itself as a service object.
2. The Service is a class that provides some useful business logic.
3. The Proxy class has a reference field that points to a service object. After the proxy finishes its processing (e.g., lazy initialization, logging, access control, caching, etc.), it passes the request to the service object.
4. The Client should work with both services and proxies via the same interface. This way you can pass a proxy into any code that expects a service object.

## Examples

### Protection Proxy

It is similar to a Decorator pattern, but it is not the same. It is a Proxy pattern because it controls access to the ‎`Car` based on the driver’s age. The Decorator pattern would be used to add new behavior, not to restrict or control access.

```go
package main

import "fmt"

type Drive interface {
 Drive()
}

type Car struct{}

func (c *Car) Drive() {
 fmt.Println("Car is being driven")
}

type Driver struct {
 Age int
}

type CarProxy struct {
 car    Car
 driver *Driver
}

func (c *CarProxy) Drive() {
 if c.driver.Age >= 16 {
  c.car.Drive()
 } else {
  fmt.Println("Driver too young!")
 }
}

func NewCarProxy(driver *Driver) *CarProxy {
 return &CarProxy{Car{}, driver}
}

func main() {
 car := NewCarProxy(&Driver{18})
 car.Drive()
}
```

### Virtual Proxy

This example shows how a Virtual Proxy can defer loading the image only when it needs to be ready.

```go
package main

import "fmt"

type Image interface {
 Draw()
}

type Bitmap struct {
 filename string
}

func NewBitmap(filename string) *Bitmap {
 fmt.Println("Loading image from", filename)
 return &Bitmap{filename}
}

func (b *Bitmap) Draw() {
 fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
 fmt.Println("About to draw the image")
 image.Draw()
 fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
 filename string
 bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
 return &LazyBitmap{filename: filename}
}

func (l *LazyBitmap) Draw() {

 if l.bitmap == nil {
  l.bitmap = NewBitmap(l.filename)
 }
 l.bitmap.Draw()
}

func main() {
 bmp := NewLazyBitmap("demo.png")
 DrawImage(bmp)
}
```
