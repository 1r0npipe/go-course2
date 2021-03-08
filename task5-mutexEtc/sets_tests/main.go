//Протестируйте производительность операций чтения и записи
//на множестве действительных чисел, безопасность которого
//обеспечивается sync.Mutex и sync.RWMutex для разных вариантов
//использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
package main

import (
	"sync"
)

// SetM is the example of working wiht conccurency
type SetM struct {
	sync.Mutex
	maps map[float64]struct{}
}

// NewSetM defines new set for Mutex version
func NewSetM() *SetM {
	return &SetM{
		maps: map[float64]struct{}{},
	}
}

// AddM is the function with adding to Mutex set
func (s *SetM) AddM(f float64) {
	s.Lock()
	s.maps[f] = struct{}{}
	s.Unlock()
}

// HasM is the function of checking the element into Mutex set
func (s *SetM) HasM(f float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.maps[f]
	return ok
}

// SetRWM is another example of working with RWMutex
type SetRWM struct {
	sync.RWMutex
	maps map[float64]struct{}
}

// NewSetRWM defines new set for Mutex version
func NewSetRWM() *SetRWM {
	return &SetRWM{
		maps: map[float64]struct{}{},
	}
}

// AddRWM is the function with adding to Mutex set
func (s *SetRWM) AddRWM(f float64) {
	s.Lock()
	s.maps[f] = struct{}{}
	s.Unlock()
}

// HasRWM is the function of checking the element into Mutex set
func (s *SetRWM) HasRWM(f float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.maps[f]
	return ok
}

