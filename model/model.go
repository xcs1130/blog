package model

//模型
import "gorm.io/gorm"

type User struct { //用户相关
	gorm.Model
	Username string `json:"username"`
	Password string `json:"passowrd"`
}

type Post struct { //博客相关
	gorm.Model
	Title   string
	Context string `gorm:"type:text"`
	Tag     string
}
