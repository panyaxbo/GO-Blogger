package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)
	log.Println("STarting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("AN Error occured Starting hTTP listen")
		log.Println("err: " + err.Error())
	}
}
