package domain

import (
	"net/http"

	"github.com/ilmsg/thawing/model"
)

type HandlerIUserContact interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type ServiceIUserContact interface {
	Create(contact *model.Contact) error
	List(page, pageSize int) (contacts []*model.Contact, err error)
	Delete(id uint) error
}

type RepoIUserContact interface {
	Create(*model.Contact) error
	List(page, pageSize int) (contacts []*model.Contact, err error)
	Delete(id uint) error
}
