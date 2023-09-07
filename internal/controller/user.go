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

func (uc UserController) List() func(c *gin.Context) {
	return func(c *gin.Context) {
		user_list.Handle(c)
	}
}

func (uc UserController) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		user_create_service.Handle(c)
	}
}

func (uc UserController) Update() func(c *gin.Context) {
	return func(c *gin.Context) {
		user_update_service.Handle(c)
	}
}

func (uc UserController) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		user_delete_service.Handle(c)
	}
}
