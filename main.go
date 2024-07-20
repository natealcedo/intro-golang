package main

import "fmt"

func main() {
	user := User{
		username: "Nate",
		age:      getNumber(),
	}
	fmt.Printf("User %+v\n", user)
}
