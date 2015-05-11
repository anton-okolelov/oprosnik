package oprosnik

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"oprosnik/model"
	"regexp"
	"strconv"
	"strings"
)

// главная страница
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	 model.DestroySession(w, r)
	
	renderExtended(w, "select-name.html", nil)
}

// Создаем сессию и сохраняем в сессии имя юзера (недо логин)
func SaveUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")

	session := model.GetUserSession(w, r)
	session.Name = name
	session.Save()
	http.Redirect(w, r, "/question", 302)
}


func Question(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := model.GetUserSession(w, r)
	if !session.IsLogged {
		http.Redirect(w, r, "/", 302)
	}

	question, answeredCount, totalCount := model.GetNextQuestion(session.Answers)
	s1, s2 := question.GetMixedSentenceIds()
	sentences := model.GetSentences()
	if totalCount > answeredCount && totalCount > 0 {
		data := map[string]interface{}{
			"name":      session.Name,
			"id1":       s1,
			"id2":       s2,
			"sentence1": sentences[s1],
			"sentence2": sentences[s2],
			"progressPercent": 100.0 * answeredCount / totalCount,
		}
		session.LastQuestion = question
		session.Save()
		renderExtended(w, "question.html", data)
	} else {
		model.GenerateReportForAdmin(session.Name, session.Answers)
		renderExtended(w, "okay.html", nil)
	}

}

// Сохранение ответа пользователей
// TODO обработка ошибок
func SaveAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session := model.GetUserSession(w, r)
	chosenSentenceId, _ := strconv.Atoi(r.FormValue("answer"))
	var answer model.Answer
	answer.Question = session.LastQuestion
	answer.ChosenSentenceId = chosenSentenceId
	session.Answers = append(session.Answers, answer)
	session.Save()
	http.Redirect(w, r, "/question", 302)
}

// главная странца админки
func Admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := &map[string]string{"sentences": strings.Join(model.GetSentences(), "\r\n")}
	renderExtended(w, "admin-form.html", params)
}

// Сохраняем список утверждений
func AdminSaveSentences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	text := r.FormValue("sentences")

	sentences := regexp.MustCompile(`(\r|\n)+`).Split(text, 1000)
	model.SaveSentences(sentences)
	http.Redirect(w, r, "/admin", 302)
}

func AdminCleanResults(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	model.DeleteReports()
	http.Redirect(w, r, "/admin", 302)
}

func StaticFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.FileServer(http.Dir("public")).ServeHTTP(w, r)
}

// обработка фатальных ошибок
func PanicHandler(w http.ResponseWriter, r *http.Request, error interface{}) {
	log.Println(error)
	w.Write([]byte("Unexpected Error"))
}
