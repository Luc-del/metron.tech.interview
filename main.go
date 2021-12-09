package main

import (
	"fmt"
	"interview/application"
	"os"
)

var VERSION string = "Luc"

func main() {
	if err := application.BuildCLIApp(VERSION).Run(os.Args); err != nil {
		fmt.Println("an error occurred:", err)
	}
}
