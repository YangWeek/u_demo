package logic

import (
	"u_demo/dao/mysql"
	"u_demo/models"
	"u_demo/pkg/jwt"
	"u_demo/pkg/snowflake"
)

// 注册用户
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err == mysql.ErrorUserExist {
		return err
	}
	//生成用户id
	userid := snowflake.GenID()
	// 生成用户实例
	user := &models.User{
		Username: p.Username,
		UserID:   userid,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

// 返回token err
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成jwt
	return jwt.GetToken(user.Username, user.UserID)
}
