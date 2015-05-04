package model

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewFilesystemStore("data/sessions", []byte("secretkey"))
var gorillaSession *sessions.Session = nil

type Session struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	IsLogged       bool
	Name           string
	Answers        []Answer
	LastQuestion   Question
}

func (this *Session) Save() {
	log.Println(this.Name)
	gorillaSession := getGorillaSession(this.request)
	gorillaSession.Values["name"] = this.Name
	gorillaSession.Values["answers"] = this.Answers
	gorillaSession.Values["lastQuestion"] = this.LastQuestion
	gorillaSession.Save(this.request, this.responseWriter)
}

func getGorillaSession(r *http.Request) *sessions.Session {
	if gorillaSession == nil {
		gorillaSession, _ = store.Get(r, "sid")
		gorillaSession.Options = &sessions.Options{
			Path: "/",
		}
	}

	return gorillaSession
}

func GetUserSession(responseWriter http.ResponseWriter, request *http.Request) *Session {
	gorillaSession := getGorillaSession(request)

	session := &Session{}
	session.request = request
	session.responseWriter = responseWriter
	nameValue, isLogged := gorillaSession.Values["name"]
	session.IsLogged = isLogged

	if isLogged {
		name, stringOk := nameValue.(string)
		if stringOk {
			session.Name = name
		}
	}

	answersValue, hasAnswers := gorillaSession.Values["answers"]
	if hasAnswers {
		answers, answersOk := answersValue.([]Answer)
		if answersOk {
			session.Answers = answers
		}
	}

	lastQuestionValue, hasLastQuestion := gorillaSession.Values["lastQuestion"]
	if hasLastQuestion {
		lastQuestion, lastQuestionOk := lastQuestionValue.(Question)
		if lastQuestionOk {
			session.LastQuestion = lastQuestion
		}
	}

	return session
}
