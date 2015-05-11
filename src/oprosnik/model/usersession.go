package model

// Работа с сессиями.
// сделана структура-прослойка Session, чтобы не делать type assertion каждый раз
// из универcального типа interface{} (в либе gorilla/sessions все хранится в map[string]interface{})

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	gorillaSession := getGorillaSession(this.request, false)
	gorillaSession.Values["name"] = this.Name
	gorillaSession.Values["answers"] = this.Answers
	gorillaSession.Values["lastQuestion"] = this.LastQuestion
	err := gorillaSession.Save(this.request, this.responseWriter)
	if err != nil {
		panic(err)
	}
}

func getGorillaSession(r *http.Request, isNew bool) *sessions.Session {
	var maxAge int
	if isNew {
		maxAge = -1
	} else {
		maxAge = 0
	}
	gorillaSession, _ := store.Get(r, "sid")
	gorillaSession.Options = &sessions.Options{
		Path:   "/",
		MaxAge: maxAge, // сессионная кука
	}

	return gorillaSession
}

func DestroySession(responseWriter http.ResponseWriter, request *http.Request) {
	gorillaSession := getGorillaSession(request, true)
	err := gorillaSession.Save(request, responseWriter)
	if err != nil {
		panic(err)
	}
}

func deleteOldFiles(path string, f os.FileInfo, err error) (e error) {
	maxTime := time.Now().Add(-3600 * time.Second)
	if strings.HasPrefix(f.Name(), "session_") && f.ModTime().Before(maxTime) {
		os.Remove(path)
	}
	return
}

func DeleteOldSessions() {
	path := "data/sessions"
	filepath.Walk(path, deleteOldFiles)
}

func SessionGarbageCollector() {
	c := time.Tick(time.Minute)
	for _ = range c {
		DeleteOldSessions()
	}
}

// берем сессию из библиотеки gorillaSession и в итоге получаем нашу прослойку Session уже с четкими
// типами, а не универсальными
func GetUserSession(responseWriter http.ResponseWriter, request *http.Request) *Session {
	gorillaSession := getGorillaSession(request, false)
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
