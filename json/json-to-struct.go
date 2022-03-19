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

	var dataUser User

	var err = json.Unmarshal(jsonData, &dataUser)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("First Name :", dataUser.FirstName)
    fmt.Println("Last Name :", dataUser.LastName)
    fmt.Println("age  :", dataUser.Age)
}