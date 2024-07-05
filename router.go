package main

import (
	"fakeApi/personGen"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ApiServer struct {
	Addr string
}

func NewApiServer(host, port string) *ApiServer {
	return &ApiServer{Addr: host + ":" + port}
}

func (a *ApiServer) Run() error {
	router := http.NewServeMux()
	server := &http.Server{Addr: a.Addr, Handler: corsMiddleware(router)}

	router.HandleFunc("GET /external", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "FAKE API SERVER")
	})

	personGen.AddRoutes(router)

	log.Printf("Starting FAKE API SERVER on %s", a.Addr)
	return server.ListenAndServe()
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
