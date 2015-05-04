package model

type Question struct {
	Sentence1Id int
	Sentence2Id int
}

type Answer struct {
	Question
	ChosenSentenceId int
}

func (this Question) isEqualTo(q Question) bool {
	return this.Sentence1Id == q.Sentence1Id && this.Sentence2Id == q.Sentence2Id ||
		this.Sentence1Id == q.Sentence2Id && this.Sentence2Id == q.Sentence1Id

}
