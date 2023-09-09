package controller

import (
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/list/user_list"
	"golang_rest_api/internal/service/user"
)

type UserController struct {
}

func (uc UserController) List(c *gin.Context) {
	user_list.Handle(c)
}

func (uc UserController) Create(c *gin.Context) {
	user.NewCreateService().Handle(c)
}

func (uc UserController) Update(c *gin.Context) {
	user.NewUpdateService().Handle(c)
}

func (uc UserController) Delete(c *gin.Context) {
	user.NewDeleteService().Handle(c)
}
