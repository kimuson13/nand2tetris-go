package main

import (
	"log"
	"os"
	"vmtranslator/process"
)

func main() {
	args := os.Args[1:]
	if err := process.Run(args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
