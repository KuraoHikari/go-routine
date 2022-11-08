package gorou_tine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T){
	go RunHelloWorld()
	fmt.Println("ups")
	time.Sleep(2 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Dispaly ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i< 100000 ; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}