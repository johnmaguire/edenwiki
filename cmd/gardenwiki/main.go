package main

import (
	"net/http"

	"github.com/johnmaguire/gardenwiki/api"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	r := api.NewRouter()

	ch := make(chan bool, 1)
	go func() {
		log.Fatal(http.ListenAndServe("0.0.0.0:3000", r))
	}()

	log.Info("Serving at 0.0.0.0:3000")
	<-ch
}
