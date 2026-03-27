package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "[SERVER] ", log.LstdFlags)
	srv := server.New(logger)
	logger.Fatal(srv.Server.ListenAndServe())

}
