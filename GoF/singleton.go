package gof

import "sync"

var lock = &sync.Mutex{}

type Singleton struct {
}

var Instance *Singleton

func Init() (*Singleton, error) {
	if Instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if Instance == nil {
			Instance = &Singleton{}
		}
	}
	return Instance, nil
}
