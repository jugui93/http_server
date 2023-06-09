package main

import (
	"context"
	"github.com/jugui93/http_server/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
 
	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8081")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("server started")
	<- serverDoneChan

	srv.Shutdown(ctx)
	log.Println("server stopped")
}