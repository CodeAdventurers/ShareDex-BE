package handler

import (
	"ShareDex-BE/src/data_model/api_model"
	"ShareDex-BE/src/service/user_service"
	"ShareDex-BE/src/util/web"
	"fmt"

	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	//// 接口校验
	userIn := &api_model.UserIn{}
	err := ctx.ShouldBindQuery(userIn)
	if err != nil {
		errMsg := fmt.Sprintf("request params parse error %s", err)
		ctx.JSON(http.StatusOK, web.FailResp(errMsg))
		return
	}

	// 调用业务逻辑处理函数
	userService := user_service.NewUserService(ctx)
	userOut := userService.GetUserInfo(userIn)

	// 返回响应参数
	ctx.JSON(http.StatusOK, web.SuccessResp(userOut))

}
