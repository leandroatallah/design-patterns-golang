package main

import "fmt"

// --- Bridge Pattern: Abstraction Hierarchy ---

// RemoteControl defines the abstraction interface for controls.
type RemoteControl interface {
	TogglePower()
	IncreaseVolume()
}

// BasicControl is a refined abstraction that works with any Device implementation.
type BasicControl struct {
	device Device // Bridge: holds a reference to the implementation interface
}

func (c *BasicControl) TogglePower() {
	c.device.SetPower(!c.device.GetPower())
}
func (c *BasicControl) IncreaseVolume() {
	c.device.SetVolume(c.device.GetVolume() + 1)
}

// AdvancedControl is another refined abstraction that works with devices supporting WiFi.
type AdvancedControl struct {
	device DeviceWithWifi // Bridge: holds a reference to the extended implementation interface
}

func (c *AdvancedControl) TogglePower() {
	c.device.SetPower(!c.device.GetPower())
}

func (c *AdvancedControl) IncreaseVolume() {
	c.device.SetVolume(c.device.GetVolume() + 1)
}
func (c *AdvancedControl) ToggleWifi() {
	c.device.SetWifi(!c.device.GetWifi())
}

// --- Bridge Pattern: Implementation Hierarchy ---

// Device defines the implementation interface for devices.
type Device interface {
	SetPower(power bool)
	GetPower() bool
	SetVolume(volume int8)
	GetVolume() int8
}

// DeviceWithWifi extends Device with WiFi capabilities.
type DeviceWithWifi interface {
	Device
	GetWifi() bool
	SetWifi(wifi bool)
}

// TV is a concrete implementation of Device.
type TV struct {
	power  bool
	volume int8
}

func (tv *TV) GetPower() bool {
	return tv.power
}
func (tv *TV) SetPower(power bool) {
	tv.power = power
}
func (tv *TV) GetVolume() int8 {
	return tv.volume
}
func (tv *TV) SetVolume(volume int8) {
	tv.volume = volume
}

// ModernTV is a concrete implementation of DeviceWithWifi.
type ModernTV struct {
	power  bool
	volume int8
	wifi   bool
}

func (tv *ModernTV) GetPower() bool {
	return tv.power
}
func (tv *ModernTV) SetPower(power bool) {
	tv.power = power
}
func (tv *ModernTV) GetVolume() int8 {
	return tv.volume
}
func (tv *ModernTV) SetVolume(volume int8) {
	tv.volume = volume
}
func (tv *ModernTV) GetWifi() bool {
	return tv.wifi
}
func (tv *ModernTV) SetWifi(wifi bool) {
	tv.wifi = wifi
}

func main() {
	sonyTV := &TV{}
	fmt.Println(sonyTV)
	basicControl := &BasicControl{sonyTV}
	basicControl.TogglePower()
	basicControl.IncreaseVolume()
	fmt.Println(sonyTV)

	androidTV := &ModernTV{}
	advControl := &AdvancedControl{androidTV}
	advControl.TogglePower()
	advControl.ToggleWifi()
	fmt.Println(androidTV)
}
