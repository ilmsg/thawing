package domain

import (
	"net/http"

	"github.com/ilmsg/thawing/model"
)

type HandlerIPage interface {
	List(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
}

type ServiceIPage interface {
	List(page, pageSize int) ([]*model.Page, error)
	Get(id uint) (*model.Page, error)
}

type RepoIPage interface {
	List() ([]*model.Page, error)
	Get(id uint) (*model.Page, error)
}
