package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type commentInterfactor struct {
	CommentRepository repository.CommentRepository
}

type CommentInterfactor interface {
	Get(c []*model.GetCommentsResponseData) ([]*model.GetCommentsResponseData, error)
	GetOne(c []*model.GetCommentsResponseData, id string) ([]*model.GetCommentsResponseData, error)
	Create(c *model.GetCommentsCreateUpdateData) (*model.GetCommentsCreateUpdateData, error)
	Update(c *model.GetCommentsCreateUpdateData, id string, userID string) (*model.GetCommentsCreateUpdateData, error)
	Delete(c []*model.Comment, id string, userID string) ([]*model.Comment, error)
}

func NewCommentInterfactor(r repository.CommentRepository) CommentInterfactor {
	return &commentInterfactor{r}
}

func (co *commentInterfactor) Get(c []*model.GetCommentsResponseData) ([]*model.GetCommentsResponseData, error) {
	c, err := co.CommentRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (co *commentInterfactor) GetOne(c []*model.GetCommentsResponseData, id string) ([]*model.GetCommentsResponseData, error) {
	c, err := co.CommentRepository.FindOne(c, id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (co *commentInterfactor) Create(c *model.GetCommentsCreateUpdateData) (*model.GetCommentsCreateUpdateData, error) {
	c, err := co.CommentRepository.Create(c)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (co *commentInterfactor) Update(c *model.GetCommentsCreateUpdateData, id string, userID string) (*model.GetCommentsCreateUpdateData, error) {
	c, err := co.CommentRepository.Update(c, id, userID)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (co *commentInterfactor) Delete(c []*model.Comment, id string, userID string) ([]*model.Comment, error) {
	c, err := co.CommentRepository.Delete(c, id, userID)
	if err != nil {
		return nil, err
	}

	return c, err
}
