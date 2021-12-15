package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


// AddCate 添加分类
func AddCate(c *gin.Context)  {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CARTNAME_USERD {
		code = errmsg.ERROR_CARTNAME_USERD
	}
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类下的文章

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// 前端获取的页数等于0时，就传入-1，不需要做Limit这个限制
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategory(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCate 编辑分类
func EditCate(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	c.ShouldBindJSON(&data)
	// 查询修改的名字是否与数据库中的重名，重名就什么都不做，否则更新数据
	code := model.CheckCate(data.Name)
	if code == errmsg.ERROR_CARTNAME_USERD {
		// c.Abort() 不执行下面的语句
		c.Abort()
		//c.JSON(http.StatusOK, gin.H{
		//	"code": code,
		//	"message": errmsg.GetErrMsg(code),
		//})
	}
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": code,
		//	"message": errmsg.GetErrMsg(code),
		//})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": errmsg.GetErrMsg(code),
	})
}
