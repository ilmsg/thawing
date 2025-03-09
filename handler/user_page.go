package handler

import (
	"net/http"

	"github.com/ilmsg/thawing/domain"
)

type UserPageHandler struct{}

func NewUserPageHandler() domain.HandlerIUserPage {
	return &UserPageHandler{}
}

func (h *UserPageHandler) List(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/page/list.html"})
}

func (h *UserPageHandler) Detail(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/page/detail.html"})
}

func (h *UserPageHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *UserPageHandler) Save(w http.ResponseWriter, r *http.Request) {

}

func (h *UserPageHandler) Update(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}

func (h *UserPageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}

func (h *UserPageHandler) Edit(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	Render(w, nil, []string{"user/page/edit.html"})
}
