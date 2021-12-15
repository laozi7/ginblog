package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb()  {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err != nil {
		fmt.Printf("connect database failed, err:%s", err)
	}
	// SingularTable设置true，则导入的表名是单数（user），否则是复数形式(users)
	// 也可以用 db.Table("自己想要的表名").CreateTable(&User{})
	db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库的最大连接数量
	db.DB().SetMaxOpenConns(100)
	// 设置连接的最大可复用数据
	db.DB().SetConnMaxLifetime(10*time.Second)

	//db.Close()
}
