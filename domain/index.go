package domain

import "net/http"

type HandlerIHome interface {
	GetIndex(w http.ResponseWriter, r *http.Request)
	GetAbout(w http.ResponseWriter, r *http.Request)
	PostContact(w http.ResponseWriter, r *http.Request)
	GetListByYear(w http.ResponseWriter, r *http.Request)
	GetListByYearMonth(w http.ResponseWriter, r *http.Request)
	GetListByYearMonthDay(w http.ResponseWriter, r *http.Request)
}
