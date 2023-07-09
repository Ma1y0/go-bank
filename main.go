package main

import (
	"log"

	"github.com/Ma1y0/go-bank/routes"
)

func main() {
    r := routes.CreateRouter()

    if err := r.Run(); err != nil {
        log.Fatal("Failed to run the router")
        return
    }

    log.Println("Server is listenig on port :8080")
}
