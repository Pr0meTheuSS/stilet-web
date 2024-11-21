package events

import (
	"sync"
)

type EventEmitter struct {
	mu        sync.RWMutex
	listeners map[string][]chan interface{}
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]chan interface{}),
	}
}

func (e *EventEmitter) Subscribe(event string, ch chan interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.listeners[event] = append(e.listeners[event], ch)
}

func (e *EventEmitter) Emit(event string, data interface{}) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	for _, ch := range e.listeners[event] {
		go func(ch chan interface{}) {
			ch <- data
		}(ch)
	}
}
