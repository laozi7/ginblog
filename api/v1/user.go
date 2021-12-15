package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// 前端获取的页数等于0时，就传入-1，不需要做Limit这个限制
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return
	}
	// 查询修改的名字是否与数据库中的重名，重名就什么都不做，否则更新数据
	code := model.CheckUser(data.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		// c.Abort() 不执行下面的语句
		c.Abort()
		//c.JSON(http.StatusOK, gin.H{
		//	"code": code,
		//	"message": errmsg.GetErrMsg(code),
		//})
	}
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": code,
		//	"message": errmsg.GetErrMsg(code),
		//})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
