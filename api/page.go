package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/johnmaguire/gardenwiki/api/data"
)

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

func (h Handlers) listPages(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.db.Pages)
}

func (h Handlers) putPage(w http.ResponseWriter, r *http.Request) {
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

	// fetch or create page
	p, ok := h.db.Pages[pageName]
	if !ok {
		p = data.Page{
			CreatedAt: time.Now(),
		}
	} else {
		p.History = append(p.History, p.PageContent)
	}

	// set new page content
	p.Body = d.Body
	p.UpdatedAt = time.Now()

	h.db.Pages[pageName] = p

	writeJSON(w, p)
}
