package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func sayHello(name string){
	fmt.Println("Hello",name)
}

func sender(ch chan int){
	ch <- 89 // send block
}

// func receiver(ch chan int){
// 	myVar:= <- ch // receive block
// 	fmt.Println("value from channel is",myVar)
// }

func main(){
	go sayHello("G1 Anshman")
	go sayHello("G2 Arpit")
	go sayHello("G3 Hardik")

	fmt.Println("Hello world!!!");

	
	ch:= make(chan int) // ch is channel
	ch2:= make(chan int)

	go sender(ch)
	go sender(ch2)
	// go receiver(ch)

	// 
	time.Sleep(time.Second * 1)

	select{
	case v:= <- ch:
		fmt.Println("Value from channel 1", v)
	case v:= <- ch2:
		fmt.Println("Value from channel 2", v)
	}

	// json.Marshal()

}
