package singleton

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

const numProcess = 100

func TestSingleton1(t *testing.T) {

	instance1 := GetInstance()
	instance2 := GetInstance()
	require.Equal(t, instance1, instance2)
}

func TestSingleton2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)

	instances := [numProcess]*Singleton{}

	for i := 0; i < numProcess; i++ {
		go func(i int) {
			instances[i] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 1; i < numProcess; i++ {
		require.Equal(t, instances[i], instances[i-1])
	}
}
