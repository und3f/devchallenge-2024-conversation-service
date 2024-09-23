package service

import (
	"log"
	"net/http"

	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/service/api"
	"github.com/gorilla/mux"
)

const (
	LISTEN_ADDR = ":8080"
)

type Service struct {
	handler http.Handler
}

func New(router *mux.Router, dao *model.Dao) *Service {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "README.md")
	})

	api.New(router, dao)

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
