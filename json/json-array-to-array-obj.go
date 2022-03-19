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
	
	jsonString := `[
		{"FirstName":"Budi","LastName":"Tabuti","Age":27},
		{"FirstName":"Udin","LastName":"Selar","Age":15}
		]`
	jsonData := []byte(jsonString)	

	
	var userArr []User
	var err = json.Unmarshal(jsonData, &userArr)	
	if err != nil {
			fmt.Println(err.Error())
			return
	}
			
	for _,value := range userArr{
		fmt.Println(value)
	}
	
}