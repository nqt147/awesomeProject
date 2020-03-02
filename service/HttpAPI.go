package service

import (
	h "github.com/awesomeProject/handler"
	m "github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *m.Router
}

func (router *Router) SetRouters() {

	router.Get("/hello", router.GetAllEmployees)
}

func (router *Router) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	h.GetAllEmployees(w, r)
}

//Restful API

func (router *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(path, f).Methods("GET")
}

func (router *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(path, f).Methods("POST")
}

func (router *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.Router.HandleFunc(path, f).Methods("PUT")
}

func (router *Router) Run(host string) {
	//log.Fatal(http.ListenAndServe(host, router.Router))
	router.Router = m.NewRouter()
	router.SetRouters()
	err := http.ListenAndServe(host, router.Router)
	if err != nil {
		panic(err)
	}
}

func (router *Router) InitializeAPIConfig() {
	router.Router = m.NewRouter()
	router.SetRouters()

}
