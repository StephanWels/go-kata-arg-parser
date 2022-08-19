package main

import (
	"fmt"
	"os"
)

type Configuration struct {
}

func main() {
	arguments := os.Args[1:]

	fmt.Println("I've parsed your input:")
	fmt.Println(ParseArguments(arguments))
}

func ParseArguments(args []string) Configuration {
	return Configuration{}
}
