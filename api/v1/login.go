package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var code int
	var token string
	c.ShouldBindJSON(&data)

	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username, data.Password)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})

}
