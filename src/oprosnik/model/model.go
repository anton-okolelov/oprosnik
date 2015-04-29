package model;

import (
	"encoding/json"
	"os"
	"path/filepath"
	"log"
)

func SaveWords(words []string) {
	wordsJson, _ := json.Marshal(words)
	f, err := os.Create("data/questions.json")
	if err != nil {
		path, _ := filepath.Abs(".")
		log.Fatal("File  " + path, err)
	}
	f.Write(wordsJson)
	f.Close()
}