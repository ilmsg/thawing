package handler

import (
	"net/http"

	"github.com/ilmsg/thawing/domain"
)

type PageHandler struct{}

func NewPageHandler() domain.HandlerIPage {
	return &PageHandler{}
}

func (h *PageHandler) Detail(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"page/detail.html"})
}

func (h *PageHandler) List(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"page/list.html"})
}
