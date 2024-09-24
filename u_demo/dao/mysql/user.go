package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"u_demo/models"
)

const (
	secret = "yang"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

//func CheckUserExist(username string) (err error) {
//	//DB.AutoMigrate(&models.User{})
//
//	user := new(models.User)
//	result := DB.Where("username = ?", username).First(user)
//	if result != nil {
//		if result.Error == gorm.ErrRecordNotFound {
//			zap.L().Error("username gorm.ErrRecordNotFound", zap.Error(result.Error))
//			return ErrorUserNotExist
//		} else {
//			zap.L().Error("Error querying user", zap.Error(result.Error))
//			return ErrorUserNotExist
//		}
//	} else {
//		return ErrorUserExist
//	}
//}
//
//// 插入
//func InsertUser(user *models.User) (err error) {
//	DB.AutoMigrate(&models.User{})
//	user.Password = encryptPassword(user.Password)
//	result := DB.Create(user)
//	if result.Error != nil {
//		zap.L().Error("InsertUser failed", zap.Error(result.Error))
//	}
//	return nil
//}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// 加密
func encryptPassword(password string) string {
	h := md5.New()            // 创建一个新的 MD5 哈希对象
	h.Write([]byte(secret))   // 写入 secret（盐）到哈希对象
	h.Write([]byte(password)) // 写入 password 到哈希对象
	// 计算哈希值并转换为字节切片
	hashedBytes := h.Sum(nil)
	return hex.EncodeToString(h.Sum([]byte(hashedBytes)))
}

//func Login(user *models.User) (err error) {
//	oPassworld := user.Password
//	//us := new(models.User)
//	// err = DB.Where("username= ?",user.Username).First(us)
//	return nil
//}

func Login(user *models.User) (err error) {
	oPassworld := user.Password
	sqlstr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlstr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	passworld := encryptPassword(oPassworld)
	if passworld != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
