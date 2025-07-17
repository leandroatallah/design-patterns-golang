package main

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = NewTorroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = NewCounterTorroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
