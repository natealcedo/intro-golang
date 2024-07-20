package util

import "github.com/natealcedo/intro-golang/types"

func GetAge(u types.User) int {
	return u.Age
}

func GetName(u types.User) string {
	return u.Username
}
