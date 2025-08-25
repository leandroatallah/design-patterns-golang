# Observer

> Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing. (Refactoring Guru)

> The observer pattern is a software design pattern in which an object, called the subject (also known as event source or event stream), maintains a list of its dependents, called observers (also known as event sinks), and automatically notifies them of any state changes, typically by calling one of their methods. (Wikipedia)

## Structure

### Observer

…

### Observable

…

## Examples

### Basic example

```go
package main

import (
 "fmt"
 "time"
)

// Business object
type Editor struct {
 events *EventManager
}

func (e *Editor) openFile(filename string) {
 event := Event{
  Filename:  filename,
  Type:      OPEN,
  Timestamp: time.Now().UTC(),
 }
 e.events.notify(event)
}

func (e *Editor) saveFile(filename string) {
 event := Event{
  Filename:  filename,
  Type:      SAVE,
  Timestamp: time.Now().UTC(),
 }
 e.events.notify(event)
}

type EventType int

const (
 OPEN EventType = iota
 SAVE
)

type Event struct {
 Type      EventType
 Filename  string
 Timestamp time.Time
}

// Subject (Publisher)
type EventManager struct {
 listeners map[EventType][]EventListener
}

func (e *EventManager) subscribe(eventType EventType, listener EventListener) {
 e.listeners[eventType] = append(e.listeners[eventType], listener)
}
func (e *EventManager) unsubscribe(eventType EventType, listener EventListener) {
 if e.listeners[eventType] != nil {
  for i, item := range e.listeners[eventType] {
   if item == listener {
    e.listeners[eventType] = append(e.listeners[eventType][:i], e.listeners[eventType][i+1:]...)
    break
   }
  }
 }
}

func (e *EventManager) notify(event Event) {
 if e.listeners[event.Type] != nil {
  for _, item := range e.listeners[event.Type] {
   item.update(event)
  }
 }
}

// Subscriber
type EventListener interface {
 update(event Event)
}

// Concrete subscribers
type EmailAlertListener struct{}

func (e *EmailAlertListener) update(event Event) {
 var action string
 if event.Type == OPEN {
  action = "open file"
 }
 if event.Type == SAVE {
  action = "save file"
 }
 fmt.Println(event.Timestamp, "[EMAIL]", action, event.Filename)
}

type LoggingListener struct {
}

func (l *LoggingListener) update(event Event) {
 var action string
 if event.Type == OPEN {
  action = "open file"
 }
 if event.Type == SAVE {
  action = "save file"
 }
 fmt.Println(event.Timestamp, "[LOGGER]", action, event.Filename)
}

func main() {
 eventManager := &EventManager{listeners: make(map[EventType][]EventListener)}
 editor := &Editor{events: eventManager}

 emailListener := &EmailAlertListener{}
 loggerListerner := &LoggingListener{}

 editor.events.subscribe(OPEN, emailListener)
 editor.events.subscribe(SAVE, emailListener)

 editor.events.subscribe(OPEN, loggerListerner)
 editor.events.subscribe(SAVE, loggerListerner)

 editor.openFile("example.html")
 editor.saveFile("example.html")
 editor.events.unsubscribe(SAVE, emailListener)
 editor.openFile("another.html")
 editor.saveFile("another.html")
}
```

### Property observer

```go
package main

import (
 "container/list"
 "fmt"
)

type Observable struct {
 subscriptions *list.List
}

func (o *Observable) Subscribe(x Observer) {
 o.subscriptions.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
 for z := o.subscriptions.Front(); z != nil; z = z.Next() {
  // if z.Value == x {
  if z.Value.(Observer) == x {
   o.subscriptions.Remove(z)
  }
 }
}

func (o *Observable) Fire(data interface{}) {
 for z := o.subscriptions.Front(); z != nil; z = z.Next() {
  z.Value.(Observer).Notify(data)
 }
}

type Observer interface {
 Notify(data interface{})
}

type PropertyChange struct {
 Name  string
 Value interface{}
}

type Person struct {
 Observable
 name string
 age  int
}

func (p *Person) Age() int {
 return p.age
}

func (p *Person) SetAge(age int) {
 if age == p.age {
  return
 }
 p.age = age
 p.Fire(PropertyChange{"Age", p.age})
}

func NewPerson(name string, age int) *Person {
 return &Person{
  Observable: Observable{new(list.List)},
  name:       name,
  age:        age,
 }
}

type TrafficManagement struct {
 o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
 if pc, ok := data.(PropertyChange); ok {
  if pc.Value.(int) >= 18 {
   fmt.Println("Congrats, you can drive now!")
   t.o.Unsubscribe(t)
  }
 }
}

func main() {
 p := NewPerson("Bruce", 15)
 t := &TrafficManagement{p.Observable}
 p.Subscribe(t)

 for i := 15; i <= 20; i++ {
  fmt.Println("Setting the age to", i)
  p.SetAge(i)
 }
}
```
