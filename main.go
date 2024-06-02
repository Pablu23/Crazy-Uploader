package main

import (
	"github.com/pablu23/uploader/server"
	// "github.com/rs/zerolog"
)

func main() {
	s := server.New()
	s.RegisterRoutes()
	s.Start()
}
