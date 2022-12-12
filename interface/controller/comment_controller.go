package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"log"
	"net/http"
	"strconv"
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
	var co []*model.GetCommentsResponseData

	co, err := cc.commentInterfactor.Get(co)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, co, "  ")
}

func (cc *commentController) GetOneComment(c Context) error {
	var co []*model.GetCommentsResponseData
	id := c.Param("id")

	co, err := cc.commentInterfactor.GetOne(co, id)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, co, "  ")
}

func (cc *commentController) CreateComment(c Context) error {
	var comment model.GetCommentsCreateUpdateData

	err := c.Bind(&comment)
	if err != nil {
		log.Print(err)
	}
	comment.UserID = getUserID(c)

	co, err := cc.commentInterfactor.Create(&comment)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Status 201. Comment №" + strconv.Itoa(int(co.ID)) + " created", "  ")
}

func (cc *commentController) UpdateComment(c Context) error {
	var comment model.GetCommentsCreateUpdateData

	err := c.Bind(&comment)
	if err != nil {
		log.Print(err)
	}

	userID := getUserID(c)
	id := c.Param("id")

	err = cc.commentInterfactor.Update(&comment, id, userID)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Status 201. Comment №" + id + " updated", "  ")
}

func (cc *commentController) DeleteComment(c Context) error {
	var co []*model.Comment

	userID := getUserID(c)
	id := c.Param("id")

	co, err := cc.commentInterfactor.Delete(co, id, userID)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Status 200. Comment №"+id+" deleted", "  ")
}
