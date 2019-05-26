package main

import (
	"fmt"
)

func main() {
	logDir, commandName, args := parser()
	fmt.Println(logDir)
	fmt.Println(commandName)
	fmt.Println(args)
}