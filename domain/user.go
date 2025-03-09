package domain

import "net/http"

type HandlerIUser interface {
	GetRegister(w http.ResponseWriter, r *http.Request)
	PostRegister(w http.ResponseWriter, r *http.Request)
	GetLogin(w http.ResponseWriter, r *http.Request)
	PostLogin(w http.ResponseWriter, r *http.Request)
	GetProfile(w http.ResponseWriter, r *http.Request)
	PostProfile(w http.ResponseWriter, r *http.Request)
	GetLogout(w http.ResponseWriter, r *http.Request)
}

type ServiceIUser interface {
}

type RepoIUser interface {
}
