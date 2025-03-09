package domain

import (
	"net/http"

	"github.com/ilmsg/thawing/model"
)

type HandlerIPost interface {
	List(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
}

type ServiceIPost interface {
	List(page, pageSize int) ([]*model.Post, error)
	Get(id uint) (*model.Post, error)
}

type RepoIPost interface {
	List() ([]*model.Post, error)
	Get(id uint) (*model.Post, error)
}
