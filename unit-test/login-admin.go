package main

type User struct{
	Username string
	Password string
}

func (user User) LoginAdmin()bool{
	if(user.Username == "admin" && user.Password == "admin"){
		return true
	}else{
		return false
	}
}
