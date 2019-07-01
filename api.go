package daikin

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"net/http"
)

type ACConfig struct {
	On bool
	Temp int
	Mode int
	Fan int
	Swing bool
}

func InitApi() {
	r := chi.NewRouter()

	r.Use(middleware.SetHeader("content-type", "application/json"))

	r.Put("/ac/livingroom", UpdateACSettings)
	r.Get("/ac/livingroom", GetACSettings)
}

func UpdateACSettings(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = Write(string(b))

	if err != nil { // TODO: improve error handling
		w.WriteHeader(500)
		return
	}

	jsonRs, err := Read()

	if err != nil { // TODO: improve error handling
		w.WriteHeader(500)
		return
	}

	w.Write(jsonRs)
	w.WriteHeader(200)
}

func GetACSettings(w http.ResponseWriter, r *http.Request) {
	jsonRs, err := Read()

	if err != nil { // TODO: improve error handling
		w.WriteHeader(500)
		return
	}

	w.Write(jsonRs)
	w.WriteHeader(200)
}