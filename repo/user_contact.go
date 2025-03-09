package repo

import (
	"github.com/ilmsg/thawing/domain"
	"github.com/ilmsg/thawing/model"
	"gorm.io/gorm"
)

type RepoUserContact struct {
	db *gorm.DB
}

func (repo *RepoUserContact) Create(contact *model.Contact) error {
	return repo.db.Create(&contact).Error
}

func (repo *RepoUserContact) Delete(id uint) error {
	var contact model.Contact
	if err := repo.db.First(&contact, id).Error; err != nil {
		return err
	}

	if err := repo.db.Delete(&contact).Error; err != nil {
		return err
	}
	return nil
}

func (repo *RepoUserContact) List(page int, pageSize int) ([]*model.Contact, error) {
	var contacts []*model.Contact
	err := repo.db.Scopes(model.Pagination(page, pageSize)).Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func NewRepoUserContact(db *gorm.DB) domain.RepoIUserContact {
	return &RepoUserContact{db}
}
