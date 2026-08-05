package main

import (
	"fmt"
	"strings"
)


func main(){
	firstName := "Aman"
	lastName := "Kar"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}