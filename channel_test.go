package gorou_tine

import (
	"fmt"
	"strconv"
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
	channel := make(chan string, 3)
	//defer close(channel)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}
	fmt.Println("selesai")
}
func TestSelectChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for{
		select {
			case data := <-channel1:
				fmt.Println("Data dari Channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari Channel 2", data)
				counter++
			
		}
		if counter == 2 {
			break
		}
	}
}

func TestSelectChannelDefault(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for{
		select {
			case data := <-channel1:
				fmt.Println("Data dari Channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari Channel 2", data)
				counter++
			default: 
				fmt.Println("Tunggu datanya ya jancok")
		}
		if counter == 2 {
			break
		}
	}
}
