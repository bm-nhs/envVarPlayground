package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error")
	}
	messageToPrint := os.Getenv("messageToPrint")
	println(messageToPrint)
}
