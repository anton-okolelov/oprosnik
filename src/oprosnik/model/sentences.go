package model

import (
	"encoding/json"
	"io/ioutil"
	"oprosnik/errors"
	"os"
	"strings"
)

var fileName = "data/sentences.json"

// сохраняем утверждения в json-файл
func SaveSentences(sentences []string) {
	nonEmptySentences := []string{}
	for _, s := range sentences {
		s = strings.TrimSpace(s)
		if s != "" {
			nonEmptySentences = append(nonEmptySentences, s)
		}
	}
	sentencesJson, err := json.Marshal(sentences)
	errors.PanicWhenError(err)
	f, err := os.Create(fileName)
	errors.PanicWhenError(err)
	f.Write(sentencesJson)
	f.Close()
}

// вытягиваем утверждения из json-файла
// TODO обработка ошибок
func GetSentences() []string {
	contents, _ := ioutil.ReadFile(fileName)
	var result []string
	_ = json.Unmarshal(contents, &result)
	return result
}
