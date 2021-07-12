package main

import (
	"net/http"
	"os"

	"github.com/johnmaguire/edenwiki/api"
	"github.com/johnmaguire/edenwiki/git"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const wikiPath = "./wiki"

func main() {
	err := os.RemoveAll(wikiPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to remove old wiki repository")
	}

	wiki, err := git.CreateWiki(wikiPath)
	r := api.NewRouter(wiki)

	ch := make(chan bool, 1)
	go func() {
		log.Fatal(http.ListenAndServe("0.0.0.0:3000", r))
	}()

	log.Info("Serving at 0.0.0.0:3000")
	<-ch
}
