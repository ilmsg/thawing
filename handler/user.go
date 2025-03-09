package handler

import (
	"net/http"

	"github.com/ilmsg/thawing/domain"
)

type UserHandler struct{}

func NewUserHandler() domain.HandlerIUser {
	return &UserHandler{}
}

func (h *UserHandler) GetRegister(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/register.html"})
}

func (h *UserHandler) PostRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/login.html"})
}

func (h *UserHandler) PostLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"user/profile.html"})
}

func (h *UserHandler) PostProfile(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) GetLogout(w http.ResponseWriter, r *http.Request) {

}
