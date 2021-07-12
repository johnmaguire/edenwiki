package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/johnmaguire/edenwiki/api/data"
	"github.com/johnmaguire/edenwiki/git"
)

type m map[string]interface{}

type putPageRequest struct {
	Body string `json:"body"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, message string) {
	e := errorResponse{message}
	buf, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(buf)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Write(buf)
}

func (h handlers) listPages(w http.ResponseWriter, r *http.Request) {
	pages, err := h.wiki.ListPages()
	if err != nil {
		writeError(w, "failed to list pages")
		return
	}

	writeJSON(w, data.PageList{Pages: pages})
}

func (h handlers) getPage(w http.ResponseWriter, r *http.Request) {
	pageName := chi.URLParam(r, "pageName")

	body, err := h.wiki.GetPage(pageName)
	switch {
	case errors.Is(err, git.ErrPageNotExists):
		http.NotFound(w, r)
		return
	case err != nil:
		panic(err)
	}

	writeJSON(w, data.Page{Body: string(body)})
}

func (h handlers) putPage(w http.ResponseWriter, r *http.Request) {
	pageName := chi.URLParam(r, "pageName")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	d := putPageRequest{}
	if err = json.Unmarshal(body, &d); err != nil {
		writeError(w, "invalid response body")
		return
	}

	if d.Body == "" {
		writeError(w, "no body provided")
		return
	}

	h.wiki.SetPage(pageName, []byte(d.Body))

	writeJSON(w, data.Page{Body: d.Body})
}
