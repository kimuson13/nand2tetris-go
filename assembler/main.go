package main

import (
	"assembler/process"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if err := process.Run(args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
