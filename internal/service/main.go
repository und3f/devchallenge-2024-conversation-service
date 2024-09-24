package service

import (
	"log"
	"net/http"

	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/service/api/call"
	"devchallenge.it/conversation/internal/service/api/category"
	"github.com/gorilla/mux"
)

const (
	LISTEN_ADDR = ":8080"
)

type Service struct {
	handler http.Handler
}

func New(router *mux.Router, dao *model.Dao, whisperUrl string) *Service {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "README.md")
	})

	apiRouter := router.PathPrefix("/api").Subrouter()
	category.Mount(apiRouter, dao)
	call.Mount(apiRouter, dao, whisperUrl)

	return &Service{
		wrapWithMiddleware(router),
	}
}

func (service *Service) Run() {
	log.Printf("Starting webserver at %q", LISTEN_ADDR)
	if err := http.ListenAndServe(LISTEN_ADDR, service.handler); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
