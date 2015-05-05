package model

import (
	"math/rand"
)

// вопрос состоит в том, какое из двух утверждений нравится больше
type Question struct {
	Sentence1Id int
	Sentence2Id int
}

type Answer struct {
	Question
	ChosenSentenceId int
}

// возвращает утверждения вопроса в случайном порядке
func (this Question) GetMixedSentenceIds() (int, int) {
	if rand.Intn(2) == 0 {
		return this.Sentence1Id, this.Sentence2Id
	} else {
		return this.Sentence2Id, this.Sentence1Id
	}
}

func (this Question) isEqualTo(q Question) bool {
	return this.Sentence1Id == q.Sentence1Id && this.Sentence2Id == q.Sentence2Id ||
		this.Sentence1Id == q.Sentence2Id && this.Sentence2Id == q.Sentence1Id

}

// проверяет вопрос в массиве вопросов или ответов
func (this Question) inArray(array interface{}) bool {
	var questions []Question
	switch v := array.(type) {
	case []Question:
		questions = v
	case []Answer:
		for _, answer := range v {
			questions = append(questions, answer.Question)
		}
	}
	for _, question := range questions {
		if this.isEqualTo(question) {
			return true
		}
	}
	return false
}
