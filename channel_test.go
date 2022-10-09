package gogoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	// defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fatih Khoiri"
		fmt.Println("Data telah terkirim")
	}()

	data := <-channel

	fmt.Println(data)
	time.Sleep(5 * time.Second)

	close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Fatih khoiri"
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)
	time.Sleep(5 * time.Second)

	close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Fatih khoiri"
}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)

	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)
	// channel <- "Fatih"
	// channel <- "Khoiri"
	// channel <- "Andi" error

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	go func() {
		channel <- "Fatih"
		channel <- "Khoiri"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima ", data)
	}
	fmt.Println("DONE")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
