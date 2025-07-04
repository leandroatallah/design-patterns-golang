package singleton

import "sync"

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton
var mu = &sync.Mutex{}

func GetInstance() Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (s *singleton) AddOne() int {
	mu.Lock()
	defer mu.Unlock()
	s.count++
	return s.count
}
