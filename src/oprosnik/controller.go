package oprosnik

import (
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"oprosnik/model"
	"regexp"
	"strconv"
	"log"
)

// главная страница
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := model.GetUserSession(w, r)
	if session.IsLogged {
		question, allAnswered := model.GetNextQuestion(*session)
		var s1, s2 int
		if rand.Intn(2) == 0 {
			s1, s2 = question.Sentence1Id, question.Sentence2Id
		} else {
			s1, s2 = question.Sentence2Id, question.Sentence1Id
		}
		sentences := model.GetSentences()
		if !allAnswered {
			data := map[string]interface{}{
				"name":      session.Name,
				"id1":       s1,
				"id2":       s2,
				"sentence1": sentences[s1],
				"sentence2": sentences[s2],
			}
			session.LastQuestion = question
			session.Save()
			renderExtended(w, "question.html", data)
		} else {
			w.Write([]byte("okay"))
			log.Println("Answers:")
			log.Println(session.Answers)			
		}
	} else {
		renderExtended(w, "select-name.html", nil)
	}
}

// Сохранение ответа пользователей
// TODO обработка ошибок
func SaveAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session := model.GetUserSession(w, r)
	chosenSentenceId, _ := strconv.Atoi(r.FormValue("answer"))
	var answer model.Answer;
	answer.Question = session.LastQuestion
	answer.ChosenSentenceId = chosenSentenceId
	session.Answers = append(session.Answers, answer)
	session.Save()
	http.Redirect(w, r, "/", 302)
}

// главная странца админки
func Admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderExtended(w, "admin-form.html", nil)
}

// Сохраняем список утверждений
func AdminSaveSentences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	text := r.FormValue("words")

	words := regexp.MustCompile("\r\n").Split(text, 1000)
	model.SaveSentences(words)
	http.Redirect(w, r, "/admin", 302)
}

func SaveUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")

	session := model.GetUserSession(w, r)
	session.Name = name
	session.Save()
	http.Redirect(w, r, "/", 302)
}
