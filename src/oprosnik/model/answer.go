package model

import (
	"math/rand"
)

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
