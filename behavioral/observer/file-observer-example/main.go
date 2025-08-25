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
