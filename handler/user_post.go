package handler

import (
	"net/http"

	"github.com/ilmsg/thawing/domain"
)

type UserPostHandler struct{}

func NewUserPostHandler() domain.HandlerIUserPost {
	return &UserPostHandler{}
}

func (h *UserPostHandler) List(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/page/list.html"})
}

func (h *UserPostHandler) Detail(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/page/detail.html"})
}

func (h *UserPostHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *UserPostHandler) Save(w http.ResponseWriter, r *http.Request) {

}

func (h *UserPostHandler) Update(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}

func (h *UserPostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}

func (h *UserPostHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	Render(w, nil, []string{"user/page/edit.html"})
}
