package main

import (
	_ "go-learning/chapter2/matchers"
	"go-learning/chapter2/search"
	"log"
	"os"
	"path/filepath"
)

func init() {
	//output log
	log.SetOutput(os.Stdout)
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {

	}
	log.Printf("current folder is %s\n", dir)
	search.Run("president")
}
