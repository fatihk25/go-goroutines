package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	fmt.Println("Halo")
	time.Sleep(1 * time.Second)
}

func TestRunAsync(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("DONE")
}
