package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/atedja/go-simple-http/app/http"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	done := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer Cleanup()

	go func() {
		<-c
		close(done)
	}()

	<-done
}

func Cleanup() {
	log.Println("Stopping services.")

	log.Println("Shutdown.")
}
