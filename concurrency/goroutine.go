package main

import (
	"fmt"	
	"time"
)

func numberLoop(){
	number := [6]int {1,2,3,4,5,6}
	for i := 0; i < len(number); i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(number[i])
	}
}


func charLoop(){
	char := [6]string {"a","b","c","d","e","f"}
	for i := 0; i < len(char); i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(char[i])
	}
}

func main(){
	
	go numberLoop()
	go charLoop()
	fmt.Scanln()
	fmt.Println("Selesai")

}