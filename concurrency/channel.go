package main

import (
	"fmt"	
	// "time"
)




func cetakAngka( c chan int) {	
		fmt.Printf("Cetak Angka: %d\n",<-c)	

}
func main() {

	chann := make(chan int)
	
	for i := 0; i < 2; i++ {
		go cetakAngka(chann)				
	}
	
	for i := 0; i < 2; i++ {
			chann <- i
		}

	fmt.Scanln()


}