package utils

import (
	"math/rand"
	"sync"
	"time"
)

type SafeRandom struct {
	random *rand.Rand
	mutex  *sync.Mutex
}

func (safeRandom *SafeRandom) Float64() float64 {
	safeRandom.mutex.Lock()
	defer safeRandom.mutex.Unlock()
	return safeRandom.random.Float64()
}

func NewSafeRandom() *SafeRandom {
	return &SafeRandom{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
		mutex:  &sync.Mutex{},
	}
}
