package models

// 用户表  基于雪花算法生成用户的id
//type User struct {
//	gorm.Model
//	UserID   int64
//	Username string
//	Password string
//}

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
