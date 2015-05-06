package model

import (
	"math/rand"
)

// берем случайные два утверждения, которые еще не спрашивали, чтобы задать вопрос
// юзеру
func GetNextQuestion(userAnswers []Answer) (question Question, allAnswered bool) {

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

	if len(questions) > 0 {
		question = questions[rand.Intn(len(questions))]
		allAnswered = false
	} else {
		allAnswered = true
	}

	return
}



