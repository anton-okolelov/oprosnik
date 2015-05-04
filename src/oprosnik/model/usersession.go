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
}

func (this *Session) Save() {
	log.Println(this.Name)
	gorillaSession := getGorillaSession(this.request)
	gorillaSession.Values["name"] = this.Name
	gorillaSession.Values["answers"] = this.Answers
	gorillaSession.Save(this.request, this.responseWriter)
}

func getGorillaSession(r *http.Request) *sessions.Session {
	if gorillaSession == nil {
		gorillaSession, _ = store.Get(r, "sid")
	}

	return gorillaSession
}

func GetUserSession(responseWriter http.ResponseWriter, request *http.Request) *Session {
	gorillaSession := getGorillaSession(request)
	nameValue, isLogged := gorillaSession.Values["name"]
	session := &Session{}
	session.IsLogged = isLogged
	if isLogged {
		name, stringOk := nameValue.(string)
		if stringOk {
			session.Name = name
		}
	}
	session.request = request
	session.responseWriter = responseWriter
	answersValue, hasAnswers := gorillaSession.Values["answers"]
	if hasAnswers {
		answers, answersOk := answersValue.([]Answer)
		if answersOk {
			session.Answers = answers
		}
	}
	return session
}
