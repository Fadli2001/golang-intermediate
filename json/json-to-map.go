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
	
	jsonString := `{"FirstName":"Budi","LastName":"Tabuti","Age":27}`
	jsonData := []byte(jsonString)	

	
	var dataInterface interface{}
	var err = json.Unmarshal(jsonData, &dataInterface)	
	if err != nil {
			fmt.Println(err.Error())
			return
	}
	
		var decodedData = dataInterface.(map[string]interface{})
		fmt.Println("First Name :", decodedData["FirstName"])
		fmt.Println("Last Name  :", decodedData["LastName"])
		fmt.Println("Age  :", decodedData["Age"])
}