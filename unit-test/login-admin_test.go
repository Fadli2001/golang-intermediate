package main

import (	
	"testing"
)

var admin User = User{"admin","admin"}



func TestLoginAdmin(test *testing.T){
	expectation := true
	actual := admin.LoginAdmin()
	if actual != expectation {
		 test.Errorf("Expected %v but got %v", expectation, actual)
	}
}