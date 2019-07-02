package daikin

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type ACConfig struct {
	On    bool
	Temp  int
	Mode  int
	Fan   int
	Swing bool
}

type Api struct {
	manager *Manager
}

func NewApi(m *Manager) *Api {
	return &Api{
		manager: m,
	}
}

func (a *Api) UpdateACSettings(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request\n")

	config := &ACConfig{}
	err := json.NewDecoder(r.Body).Decode(config)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	a.manager.SerialChan <- config

	// success
	w.WriteHeader(200)
}

func (a *Api) GetACSettings(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(a.manager.State)

	if err != nil { // TODO: improve error handling
		w.WriteHeader(500)
		return
	}

	// write response
	_, err = w.Write(b)

	if err != nil { // TODO: improve error handling
		w.WriteHeader(500)
		return
	}

	// success
	w.WriteHeader(200)
}

func (a *Api) Run() chan error {
	r := chi.NewRouter()

	r.Use(middleware.SetHeader("content-type", "application/json"))

	r.Put("/ac/livingroom", a.UpdateACSettings)
	r.Get("/ac/livingroom", a.GetACSettings)

	c := make(chan error)

	c <- http.ListenAndServe(":8080", r) // TODO: configurable addr

	fmt.Printf("Started server at port 8080\n")

	return c
}
