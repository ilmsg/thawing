package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ilmsg/thawing/domain"
)

type HomeHandler struct{}

func NewHomeHandler() domain.HandlerIHome {
	return &HomeHandler{}
}

func (h *HomeHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"index.html"})
}

func (h *HomeHandler) GetAbout(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"about.html"})
}

func (h *HomeHandler) PostContact(w http.ResponseWriter, r *http.Request) {

}

func (h *HomeHandler) GetListByYear(w http.ResponseWriter, r *http.Request) {
	// year := r.URL.Query().Get("year")
	Render(w, nil, []string{"index.html"})
}

func (h *HomeHandler) GetListByYearMonth(w http.ResponseWriter, r *http.Request) {
	// year := r.URL.Query().Get("year")
	// month := r.URL.Query().Get("month")
	Render(w, nil, []string{"index.html"})
}

func (h *HomeHandler) GetListByYearMonthDay(w http.ResponseWriter, r *http.Request) {
	// year := r.URL.Query().Get("year")
	// month := r.URL.Query().Get("month")
	// day := r.URL.Query().Get("day")
	Render(w, nil, []string{"index.html"})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{filepath.Join("templates", "layout.html")}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Fatal(err)
	}
}
