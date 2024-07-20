package main

import (
	"fmt"
	"github.com/natealcedo/intro-golang/types"
)

func main() {
	user := types.User{
		Username: "Nate",
		Age:      32,
	}
	fmt.Printf("User: %+v\n", user)
}
