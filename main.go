package main

import (
	"log"

	"github.com/seblegall/gsync/cmd"
)

func main() {
	if err := cmd.NewGsyncCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
