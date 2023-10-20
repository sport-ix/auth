package main

import (
	"github.com/sport-ix/auth/internal/auth"
	"log"
)

func main() {
	srv := new(auth.Server)
	if err := srv.Run("port"); err != nil {
		log.Fatalf("error occured while running http server : %s", err.Error())
	}
}
