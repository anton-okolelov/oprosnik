package model

import (
	"io/ioutil"
	"log"
	"sort"
	"time"
	"fmt"
)

type keyval struct {
	Sentence string
	Score    int
}

type ByScore []keyval

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func GenerateReportForAdmin(name string, answers []Answer) {
	sentences := GetSentences()
	var counter = map[string]int{}
	for id, sentence := range sentences {
		counter[sentence] = 0
		for _, answer := range answers {
			if answer.ChosenSentenceId == id {
				counter[sentence]++
			}
		}
	}
	var rating = []keyval{}
	for sentence, score := range counter {
		rating = append(rating, keyval{sentence, score})
	}

	sort.Sort(ByScore(rating))

	report := "Имя: " + name + "\r\n"
	report += "--------------------\r\n"
	for _, keyval := range rating {
		report += fmt.Sprintf("%d", keyval.Score) + "\t" + keyval.Sentence + "\r\n"
	}

	log.Println(report)

	ioutil.WriteFile("data/results"+time.Now().Format(time.RFC3339)+".txt", []byte(report), 0600)
}
