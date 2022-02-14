package utils

import (
	"fmt"
)

// Function for handling errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
