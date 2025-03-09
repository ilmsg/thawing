package handler

import (
	"net/http"

	"github.com/ilmsg/thawing/domain"
)

type PostHandler struct{}

func NewPostHandler() domain.HandlerIPost {
	return &PostHandler{}
}

func (h *PostHandler) Detail(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"post/detail.html"})
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"post/list.html"})
}
