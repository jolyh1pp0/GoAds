package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type commentInterfactor struct {
	CommentRepository repository.CommentRepository
}

type CommentInterfactor interface {
	Get(c []*model.Comment) ([]*model.Comment, error)
	GetOne(c []*model.Comment, id string) ([]*model.Comment, error)
	Create(c *model.Comment) (*model.Comment, error)
	Update(c *model.Comment, id string) (*model.Comment, error)
	Delete(c []*model.Comment, id string) ([]*model.Comment, error)
}

func NewCommentInterfactor(r repository.CommentRepository) CommentInterfactor {
	return &commentInterfactor{r}
}

func (co *commentInterfactor) Get(c []*model.Comment) ([]*model.Comment, error) {
	c, err := co.CommentRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (co *commentInterfactor) GetOne(c []*model.Comment, id string) ([]*model.Comment, error) {
	c, err := co.CommentRepository.FindOne(c, id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (co *commentInterfactor) Create(c *model.Comment) (*model.Comment, error) {
	c, err := co.CommentRepository.Create(c)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (co *commentInterfactor) Update(c *model.Comment, id string) (*model.Comment, error) {
	c, err := co.CommentRepository.Update(c, id)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (co *commentInterfactor) Delete(c []*model.Comment, id string) ([]*model.Comment, error) {
	c, err := co.CommentRepository.Delete(c, id)
	if err != nil {
		return nil, err
	}

	return c, err
}
