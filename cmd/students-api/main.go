package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Khanhasan26/students-api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setups

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	//setup server
	server := http.Server{

		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("Server Started %s\n", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server")
	}

}
