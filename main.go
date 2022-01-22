package main

import (
	"log"
	"log-generator/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
