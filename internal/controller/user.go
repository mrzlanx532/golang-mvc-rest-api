package controller

import (
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/list/user_list"
	"golang_rest_api/internal/service/user_create_service"
	"golang_rest_api/internal/service/user_delete_service"
	"golang_rest_api/internal/service/user_update_service"
)

type UserController struct {
}

func (uc UserController) List(c *gin.Context) {
	user_list.Handle(c)
}

func (uc UserController) Create(c *gin.Context) {
	user_create_service.Handle(c)
}

func (uc UserController) Update(c *gin.Context) {
	user_update_service.Handle(c)
}

func (uc UserController) Delete(c *gin.Context) {
	user_delete_service.Handle(c)
}
