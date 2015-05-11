package model

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
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

	t := time.Now()
	fileName := fmt.Sprintf("%d-%02d-%02d_%02d-%02d-%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) + ".txt"
	ioutil.WriteFile("public/results/"+fileName, []byte(report), 0600)
}

func DeleteReports() {
	dirname := "public/results/"

	d, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".txt" {
				os.Remove(dirname + file.Name())
			}
		}
	}
}
