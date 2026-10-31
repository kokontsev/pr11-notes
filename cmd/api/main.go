// @title Notes API
// @version 1.0
// @description REST API for notes
// @BasePath /api/v1
package main

import (
	"example.com/pr11/internal/httpapi"
	"example.com/pr11/internal/repo"
	"example.com/pr11/internal/service"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	repository := repo.NewMemory()
	h := &httpapi.Handler{Service: service.New(repository)}
	m := h.Router()
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal((&http.Server{Addr: addr, Handler: m, ReadHeaderTimeout: 5 * time.Second}).ListenAndServe())
}
