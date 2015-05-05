package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var fileName = "data/sentences.json"

// TODO обработка ошибок
func SaveSentences(sentences []string) {
	sentencesJson, _ := json.Marshal(sentences)
	f, err := os.Create(fileName)
	if err != nil {
		path, _ := filepath.Abs(".")
		log.Fatal("File  "+path, err)
	}
	f.Write(sentencesJson)
	f.Close()
}

// TODO обработка ошибок
func GetSentences() []string {
	contents, _ := ioutil.ReadFile(fileName)
	var result []string
	_ = json.Unmarshal(contents, &result)
	return result
}
