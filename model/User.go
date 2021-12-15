package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 如果一个字段的tag加上了 binding:"required"， 但绑定时是空值，Gin会报错，也就是说必须绑定字段
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);not null" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
	// Role 权限设置
	Role int `gorm:"type:int" json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
// pageSize 一页有多少条数据
// PageName 当前页数

func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	// 查询pageSize条数据，从(pageNum - 1) * pageSize 条数据后开始查
	// 例如：当前页数是第三页，这一页要显示5条数据，那么((3-1)*5)=10条，偏移10条数据，返回第11条到第15条数据
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	// 如果有错误，就返回 nil
	// 传入接收结果集的变量只能为Struct类型或Slice类型
	// 当传入变量为Struct类型时，如果检索出来的数据为0条，会抛出ErrRecordNotFound错误
	// 当传入变量为Slice类型时，任何条件下均不会抛出ErrRecordNotFound错误
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUser 编辑用户

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ScryptPw(password string) string {
	//const salt = "salt"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func ValidatePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if !ValidatePasswords(user.Password, password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
