package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

// Category 文章分类
type Category struct {
	ID int `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(30);not null" json:"name"`
}

// GetCategory 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 查询分类下的文章

// CreateCategory 添加分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err !=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCate 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}


// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckCate 查询分类是否存在
func CheckCate(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CARTNAME_USERD
	}
	return errmsg.SUCCESS
}

