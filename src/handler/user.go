package handler

import (
	"ShareDex-BE/src/util/web"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	//// 接口校验
	//userIn := &api_models.UserIn{}
	//err := c.ShouldBindQuery(userIn)
	//if err != nil {
	//	errMsg := fmt.Sprintf("request params parse error %s", err)
	//	c.JSON(http.StatusOK, web.FailResp(errMsg))
	//	return
	//}
	//
	//// 调用业务逻辑处理函数
	//userLogic := user_logic.NewUserLogic()
	//userOut := userLogic.GetUserInfo(userIn)

	// 返回响应参数
	c.JSON(http.StatusOK, web.SuccessResp(nil))

}
