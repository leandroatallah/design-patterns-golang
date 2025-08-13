# Mediator

> Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object. (Refactoring Guru)

> The Mediator pattern defines an object that encapsulates how a set of objects interact. (Wikipedia)

## Structure

- From a component’s perspective, it all looks like a total black box. The sender doesn’t know who’ll end up handling its request, and the receiver doesn’t know who sent the request in the first place.

## Example

### Basic

```go
package main

import (
 "fmt"
)

// Mediator interface
type Mediator interface {
 notify(sender Component, event EVENT)
}

// Concrete mediator
type AuthenticationDialog struct {
 title            string
 okBtn, cancelBtn *Button
 checkbox         *Checkbox
}

func NewAuthenticationDialog(title string, okBtn, cancelBtn *Button, checkbox *Checkbox) *AuthenticationDialog {
 m := &AuthenticationDialog{
  title: title, okBtn: okBtn,
  cancelBtn: cancelBtn,
  checkbox:  checkbox,
 }
 okBtn.setDialog(m)
 cancelBtn.setDialog(m)
 checkbox.setDialog(m)
 return m
}

func (a *AuthenticationDialog) notify(sender Component, event EVENT) {
 if sender == a.okBtn && event == CLICK {
  fmt.Println("Ok Button clicked.")
 }
 if sender == a.okBtn && event == KEYPRESS {
  fmt.Println("Ok Button key pressed.")
 }
 if sender == a.cancelBtn && event == CLICK {
  fmt.Println("Cancel Button clicked.")
 }
 if sender == a.cancelBtn && event == KEYPRESS {
  fmt.Println("Cancel Button key pressed.")
 }
 if sender == a.checkbox && event == CHECK {
  fmt.Println("Checkbox checked.")
 }
}

// Events
type EVENT int

const (
 CLICK EVENT = iota
 KEYPRESS
 CHECK
)

// Abstract component
type Component interface {
 setDialog(dialog Mediator)
}

// Concrete component
type Button struct {
 dialog Mediator
}

func NewButtonComponent() *Button {
 return &Button{}
}

func (b *Button) setDialog(dialog Mediator) {
 b.dialog = dialog
}

func (b *Button) click() {
 b.dialog.notify(b, CLICK)
}

func (b *Button) keypress() {
 b.dialog.notify(b, KEYPRESS)
}

type Checkbox struct {
 dialog Mediator
}

func NewCheckboxComponent() *Checkbox {
 return &Checkbox{}
}

func (c *Checkbox) check() {
 c.dialog.notify(c, CHECK)
}

func (c *Checkbox) setDialog(dialog Mediator) {
 c.dialog = dialog
}

func main() {
 okBtn := NewButtonComponent()
 cancelBtn := NewButtonComponent()

 checkbox := NewCheckboxComponent()
 NewAuthenticationDialog("AuthenticationDialog", okBtn, cancelBtn, checkbox)

 // sandbox
 okBtn.click()
 okBtn.keypress()
 checkbox.check()
}
```

## Chat room

```go
package main

import "fmt"

type Person struct {
 Name    string
 Room    *ChatRoom
 chatLog []string
}

func NewPerson(name string) *Person {
 return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
 s := fmt.Sprintf("%s: %s", sender, message)
 fmt.Printf("[%s's chat session]: %s\n", p.Name, s)
 p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
 p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
 p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
 people []*Person
}

// Show the message to the others
func (c *ChatRoom) Broadcast(source, message string) {
 for _, p := range c.people {
  if p.Name != source {
   p.Receive(source, message)
  }
 }
}

// Show the message only yo the destination
func (c *ChatRoom) Message(source, destination, message string) {
 for _, p := range c.people {
  if p.Name == destination {
   p.Receive(source, message)
  }
 }
}

func (c *ChatRoom) Join(p *Person) {
 joinMsg := p.Name + " joins the chat"
 c.Broadcast("Room", joinMsg)

 p.Room = c
 c.people = append(c.people, p)
}

func main() {
 room := ChatRoom{}
 john := NewPerson("John")
 jane := NewPerson("Jane")

 room.Join(john)
 room.Join(jane)

 john.Say("Hi room")
 jane.Say("oh, hey John")

 simon := NewPerson("Simon")
 room.Join(simon)
 simon.Say("Hi guys!")

 jane.PrivateMessage("Simon", "glad you could join us!")
}
```

### Train station

```go
package main

import "fmt"

// train
type Train interface {
 arrive()
 depart()
 permitArrival()
}

// passenger
type PassengerTrain struct {
 mediator Mediator
}

func (g *PassengerTrain) arrive() {
 if !g.mediator.canArrive(g) {
  fmt.Println("PassengerTrain: Arrival blocked, waiting")
  return
 }
 fmt.Println("PassengerTrain: Arrived")
}
func (g *PassengerTrain) depart() {
 fmt.Println("PassengerTrain: leaving")
 g.mediator.notifyAboutDeparture()
}
func (g *PassengerTrain) permitArrival() {
 fmt.Println("PassengerTrain: Arrival permitted, arriving")
 g.arrive()
}

// freit train
type FreightTrain struct {
 mediator Mediator
}

func (g *FreightTrain) arrive() {
 if !g.mediator.canArrive(g) {
  fmt.Println("FreightTrain: Arrival blocked, waiting")
  return
 }
 fmt.Println("FreightTrain: Arrived")
}
func (g *FreightTrain) depart() {
 fmt.Println("FreightTrain: leaving")
 g.mediator.notifyAboutDeparture()
}
func (g *FreightTrain) permitArrival() {
 fmt.Println("FreightTrain: Arrival permitted, arriving")
 g.arrive()
}

// Mediator
type Mediator interface {
 canArrive(Train) bool
 notifyAboutDeparture()
}

// Statios Manager
type StationManager struct {
 isPlatformFree bool
 trainQueue     []Train
}

func NewStationManager() *StationManager {
 return &StationManager{
  isPlatformFree: true,
 }
}

func (s *StationManager) canArrive(train Train) bool {
 if s.isPlatformFree {
  s.isPlatformFree = false
  return true
 }
 s.trainQueue = append(s.trainQueue, train)
 return false
}

func (s *StationManager) notifyAboutDeparture() {
 if !s.isPlatformFree {
  s.isPlatformFree = true
 }
 if len(s.trainQueue) > 0 {
  firstTrainQueue := s.trainQueue[0]
  s.trainQueue = s.trainQueue[1:]
  firstTrainQueue.permitArrival()
 }
}

func main() {
 stationManager := NewStationManager()

 passengerTrain := &PassengerTrain{
  mediator: stationManager,
 }
 freightTrain := &FreightTrain{
  mediator: stationManager,
 }

 passengerTrain.arrive()
 freightTrain.arrive()
 passengerTrain.depart()
}
```
