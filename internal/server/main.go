package main

import (
	"log"

	"github.com/ahonculada/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8079")
	log.Fatal(srv.ListenAndServer())
}
