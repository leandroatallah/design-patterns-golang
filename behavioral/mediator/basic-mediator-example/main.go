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
