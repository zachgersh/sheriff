package main

import (
	"log"

	"github.com/zachgersh/sheriff/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
