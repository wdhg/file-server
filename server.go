package main

import (
	"fmt"
	"os"
)

func Create(file string) {
	os.Create("files/" + file)
}

func main() {
	fmt.Println("hello")
}
