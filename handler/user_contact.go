package handler

import "net/http"

type UserContactHandler struct{}

func (h *UserContactHandler) List(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/contact/list.html"})
}

func (h *UserContactHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}
