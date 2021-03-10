package singleton

import "sync"

var singleton *Singleton
var once sync.Once

type Singleton struct {
}

func GetInstance() *Singleton {

	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
