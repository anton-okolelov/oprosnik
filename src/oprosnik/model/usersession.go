package model

// Работа с сессиями.
// сделана структура-прослойка Session, чтобы не делать type assertion каждый раз
// из универcального типа interface{} (в либе gorilla/sessions все хранится в map[string]interface{})

import (
	"github.com/gorilla/sessions"
	//"log"
	"encoding/gob"
	"net/http"
)

var store = sessions.NewFilesystemStore("data/sessions", []byte("secretkey"))

func init() {
	// магия, чтобы работала сериализация в сессию
	gob.Register([]Answer{})
	gob.Register(Question{})
}

type Session struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	IsLogged       bool
	Name           string
	Answers        []Answer
	LastQuestion   Question
}

func (this *Session) Save() {
	gorillaSession := getGorillaSession(this.request)
	gorillaSession.Values["name"] = this.Name
	gorillaSession.Values["answers"] = this.Answers
	gorillaSession.Values["lastQuestion"] = this.LastQuestion
	err := gorillaSession.Save(this.request, this.responseWriter)
	if err != nil {
		panic(err)
	}
}

func getGorillaSession(r *http.Request) *sessions.Session {

	gorillaSession, _ := store.Get(r, "sid")
	gorillaSession.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 0, // сессионная кука
		HttpOnly: true,
	}

	return gorillaSession
}

// берем сессию из библиотеки gorillaSession и в итоге получаем нашу прослойку Session уже с четкими
// типами, а не универсальными
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
