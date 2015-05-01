package model

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func SaveWords(words []string) {
	wordsJson, _ := json.Marshal(words)
	f, err := os.Create("data/questions.json")
	if err != nil {
		path, _ := filepath.Abs(".")
		log.Fatal("File  "+path, err)
	}
	f.Write(wordsJson)
	f.Close()
}
