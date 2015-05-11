package model

import (
	"math/rand"
)

// берем случайные два утверждения, которые еще не спрашивали, чтобы задать вопрос
// юзеру
func GetNextQuestion(userAnswers []Answer) (question Question, answeredCount int, totalCount int) {

	sentences := GetSentences()
	var questions []Question
	for sentence1 := range sentences {
		for sentence2 := range sentences {
			if sentence1 == sentence2 {
				continue
			}

			question := Question{sentence1, sentence2}

			if !question.inArray(userAnswers) && !question.inArray(questions) {
				questions = append(questions, question)
			}
		}
	}

	answeredCount = len(userAnswers)
	totalCount = len(userAnswers) + len(questions)

	if len(questions) > 0 {
		question = questions[rand.Intn(len(questions))]
	}

	return
}
