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
