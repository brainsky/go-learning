package main

import (
	_ "go-learning/chapter2/matchers"
	"go-learning/chapter2/search"
	"log"
	"os"
)

func init() {
	//output log
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
