package controller

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"devchallenge.it/conversation/internal/controller/api/call"
	"devchallenge.it/conversation/internal/controller/api/category"
	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/services"
	"github.com/gorilla/mux"
)

const (
	LISTEN_ADDR = ":8080"
)

type Service struct {
	server *http.Server
}

func New(router *mux.Router, dao *model.Dao, servConf model.ServicesConf) *Service {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "README.md")
	})

	server := &http.Server{Addr: LISTEN_ADDR, Handler: wrapWithMiddleware(router)}
	BindQuit(server)

	services := services.CreateServicesFacade(servConf)
	apiRouter := router.PathPrefix("/api").Subrouter()

	category.Mount(apiRouter, dao)
	call.Mount(apiRouter, dao, services)

	return &Service{
		server,
	}
}

func (service *Service) Run() {
	log.Printf("Starting webserver at %q", LISTEN_ADDR)

	if err := service.server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func (service *Service) CreateServices(srvConf model.ServicesConf) services.ServicesFacade {
	return services.CreateServicesFacade(srvConf)
}

func BindQuit(server *http.Server) chan any {
	quit := make(chan any, 1)
	go func() {
		sigQuit := make(chan os.Signal, 1)
		signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)
		<-sigQuit
		quit <- nil

		if err := server.Close(); err != nil {
			log.Fatalf("HTTP close error: %v", err)
		}
	}()

	return quit
}
