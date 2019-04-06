package search

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const dataFile = "data\\data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	current, err := os.Getwd()
	dataPath := filepath.Join(current, "chapter2\\data\\data.json")
	file, err := os.Open(dataPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed

	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
