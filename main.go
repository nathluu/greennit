package main

import (
	"flag"
	"github.com/nathluu/greennit/internal/app/api"
	"github.com/nathluu/greennit/internal/pkg/http/server"
	"github.com/nathluu/greennit/internal/pkg/log"
)

func main() {
	port := flag.String("p", "8080", "Service port")

	router, err := api.NewRouter()
	if err != nil {
		log.Panicf("failed to init router: %v\n", err)
	}
	server.ListenAndServe(nil, router)
}
