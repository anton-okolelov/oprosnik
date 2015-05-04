package model

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func SaveSentences(sentences []string) {
	sentencesJson, _ := json.Marshal(sentences)
	f, err := os.Create("data/sentences.json")
	if err != nil {
		path, _ := filepath.Abs(".")
		log.Fatal("File  "+path, err)
	}
	f.Write(sentencesJson)
	f.Close()
}
