package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTic(t *testing.T) {
	ticker := time.Tick(1 * time.Second)

	for time := range ticker {
		fmt.Println(time)
	}
}
