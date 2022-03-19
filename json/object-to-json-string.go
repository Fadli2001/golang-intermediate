package main

import (
	"fmt"
	"encoding/json"
)


type User struct{
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Age int
}

func main(){
	var arrayStruct = [2]User{
		{"Dena","Dwiya", 27}, 
		{"Eka","Sura", 32},
	}
	var jsonData, err = json.Marshal(arrayStruct)
	if err != nil {
			fmt.Println(err.Error())
			return
	}
	
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}