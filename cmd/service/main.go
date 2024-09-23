package main

import (
	"devchallenge.it/conversation/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	service.New(mux.NewRouter(), service.NewDao()).Run()
}
