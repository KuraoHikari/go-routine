package gorou_tine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCannel (t *testing.T){

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Kurao Hikari"

		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel

	fmt.Println(data)
	time.Sleep(3 * time.Second)
}

func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Kurao Hikari V3"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Kurao Hikari V4"
}
func OnlyOut(channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Kurao"
		channel <- "Hikari"
		channel <- "V2"
	}()
	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(3 * time.Second)
}

func TestRangeChannel(t *testing.T){
	
}