package singleton_mode

import (
	"github.com/stretchr/testify/assert"
	"github.com/ullt/gopool"
	"runtime"
	"testing"
)

func TestGetSingle(t *testing.T) {
	as := assert.New(t)
	s1 := GetSingle()
	s2 := GetSingle()
	as.Equal(s1, s2)
}

func TestGetSingleWithGoroutines(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	as := assert.New(t)
	s := GetSingle()
	sChan := make(chan ISingleton, 100)
	pool := gopool.New(100, 100)
	for i := 0; i < 100; i++ {
		pool.WaitCount(1)
		pool.Do(func() {
			s := GetSingle()
			sChan <- s
			pool.TaskDone()
		})
	}
	pool.WaitAll()
	close(sChan)
	for singleton := range sChan {
		as.Equal(s, singleton)
	}
}
