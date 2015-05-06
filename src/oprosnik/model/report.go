package model

import (
	"sort"
	"log"
	"encoding/json"
	"io/ioutil"
	"time"
)

type keyval struct {
	Sentence string
	Score int
}

type ByScore []keyval

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func GenerateReportForAdmin(name string, answers []Answer) {
	sentences := GetSentences()
	var counter = map[string]int{}
	for id, sentence := range(sentences) {
		counter[sentence] = 0
		for _, answer := range answers {
			if (answer.ChosenSentenceId == id) {
				counter[sentence]++
			}
		}
	}
	var rating = []keyval{}
	for sentence, score := range(counter) {
		rating = append(rating, keyval{sentence, score})
	}

	sort.Sort(ByScore(rating))
	
	report := "Имя: " + name + "\r\n"
	report += "--------------------"
	for _, keyval := range(rating) {
		report += keyval.Score + "\t" + keyval.Sentence
	}
	
	log.Println(report)
	//jsonContent, err := json.Marshal(rating)
	//ioutil.WriteFile("data/results" + string(time.Now()) + ".json")
}