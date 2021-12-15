package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	// Cid 关联到 Category 的 id
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`
	// Desc 文章描述
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	// Content 文章内容
	Content string `gorm:"type:longtext" json:"content"`
	// Img 文章图片
	Img string `gorm:"type:varchar(200)" json:"img"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArts 查询分类下所有文章
func GetCateArts(id, pageSize, pageNum int) ([]Article, int) {
	var cateArtList []Article
	// preloads 预加载category表，传入分类ID，查询该分类下所有文章
	err := db.Preloads("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("id = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CAR_NOT_EXIST
	}
	return cateArtList, errmsg.SUCCESS
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preloads("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetArts 查询文章列表
func GetArts(pageSize int, pageNum int) ([]Article,int) {
	var articleList []Article
	err =db.Preloads("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["Content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
