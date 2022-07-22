package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"log"
	"net/http"
)

type commentController struct {
	commentInterfactor interfactor.CommentInterfactor
}

type CommentController interface {
	GetComments(c Context) error
	GetOneComment(c Context) error
	UpdateComment(c Context) error
	CreateComment(c Context) error
	DeleteComment(c Context) error
}

func NewCommentController(co interfactor.CommentInterfactor) CommentController {
	return &commentController{co}
}

func (cc *commentController) GetComments(c Context) error {
	var co []*model.Comment

	co, err := cc.commentInterfactor.Get(co)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, co, "  ")
}

func (cc *commentController) GetOneComment(c Context) error {
	var co []*model.Comment
	id := c.Param("id")

	co, err := cc.commentInterfactor.GetOne(co, id)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, co, "  ")
}

func (cc *commentController) CreateComment(c Context) error {
	var comment model.Comment

	err := c.Bind(&comment)
	if err != nil {
		log.Print(err)
	}

	co, err := cc.commentInterfactor.Create(&comment)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, co, "  ")
}

func (cc *commentController) UpdateComment(c Context) error {
	var comment model.Comment

	err := c.Bind(&comment)
	if err != nil {
		log.Print(err)
	}

	id := c.Param("id")

	co, err := cc.commentInterfactor.Update(&comment, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, co, "  ")
}

func (cc *commentController) DeleteComment(c Context) error {
	var co []*model.Comment

	id := c.Param("id")

	co, err := cc.commentInterfactor.Delete(co, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Comment "+id+" deleted", "  ")
}
