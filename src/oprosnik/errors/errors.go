package errors

import (
	"net/http"
	"log"
)

// обработка фатальных ошибок
func PanicHandler(w http.ResponseWriter, r *http.Request, error interface{}) {
	log.Println(error)
	w.Write([]byte("Unexpected Error"))
}

func PanicWhenError(err error) {
	if (err != nil) {
		panic(err)
	}
}