package singleton

import "sync"

// this pattern is used when only one instance of a struct should exist
// it is applicable for DB or Logger instances

// single is the struct which should be instantiated only once
type single struct{}

// the instance should be a pointer
var singleInstance *single

// METHOD 1: using lock/unlock operations
var lock = &sync.Mutex{}

// getInstance returns the existing instance of single
// the instance is created the first time getInstance is called
func getInstance() *single {
	//  the first check for nil instance is to prevent the cost of locking/unlocking
	// the second check is necessary to prevent two goroutines from creating two different instances concurrently
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
		}
	}
	return singleInstance
}

// METHOD 2: using once.Do()
var once sync.Once

func getInstance2() *single {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = &single{}
		})
	}
	return singleInstance
}
