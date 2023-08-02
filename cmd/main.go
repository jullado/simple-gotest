package main

import (
	"example-gotest/login"
	"fmt"
)

func main() {
	result := login.Login("", "")
	fmt.Println(result)
}
