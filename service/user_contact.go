package service

import (
	"github.com/ilmsg/thawing/domain"
	"github.com/ilmsg/thawing/model"
)

type ServiceUserContact struct {
	repo domain.RepoIUserContact
}

func (s *ServiceUserContact) Create(contact *model.Contact) error {
	return s.repo.Create(contact)
}

func (s *ServiceUserContact) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *ServiceUserContact) List(page int, pageSize int) ([]*model.Contact, error) {
	return s.repo.List(page, pageSize)
}

func NewServiceUserContact(repo domain.RepoIUserContact) domain.ServiceIUserContact {
	return &ServiceUserContact{repo}
}
