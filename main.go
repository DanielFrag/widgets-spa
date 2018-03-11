package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/DanielFrag/widgets-spa-rv/router"
	"github.com/DanielFrag/widgets-spa-rv/infra"
)

func main() {
	defer mainRecover()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		clean()
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	infra.StartDB()
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":" + port, r))
}

func clean() {
	stopDb()
	os.Exit(1)
}

func stopDb() {
	infra.StopDB()
	fmt.Println("DB closed")
}

func mainRecover() {
	rec := recover()
	if rec != nil {
		log.Println(rec)
		clean()
	}
}
