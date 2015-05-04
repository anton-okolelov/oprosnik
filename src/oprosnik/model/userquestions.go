package model

import (
	//"log"
	"math/rand"
)

func GetNextQuestion(session Session) (question Question, allAnswered bool) {

	sentences := GetSentences()
	var questions []Question
	for sentence1 := range sentences {
		for sentence2 := range sentences {
			if sentence1 == sentence2 {
				continue
			}

			question := Question{sentence1, sentence2}

			if alreadyAnswered(question, session) {
				continue
			}

			if checkIfAlreadyAdded(questions, question) {
				continue
			}

			questions = append(questions, question)
		}
	}

	if len(questions) > 0 {
		question = questions[rand.Intn(len(questions))]
		allAnswered = false
	} else {
		allAnswered = true
	}

	return
}

func alreadyAnswered(question Question, session Session) bool {
	answers := session.Answers

	for _, answer := range answers {
		if answer.Question.isEqualTo(question) {
			return true
		}
	}
	return false
}

func checkIfAlreadyAdded(questions []Question, question Question) bool {
	for _, q := range questions {
		if q.isEqualTo(question) {
			return true
		}
	}
	return false
}
