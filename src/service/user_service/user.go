package user_service

import (
	"ShareDex-BE/src/data_model/api_model"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	reqContext *gin.Context
}

func (us *UserService) GetUserInfo(userIn *api_model.UserIn) map[string]string {
	return map[string]string{
		"name": "hui",
	}
}

func NewUserService(ctx *gin.Context) *UserService {
	return &UserService{reqContext: ctx}
}
