package services

import (
	"errors"
	"footmap/app/dao"
	"footmap/app/models"
	"footmap/app/utils"
	"github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
	"strings"
	"time"
)
/**
用户登陆
 */
func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "用户名或密码不能为空", errors.New("用户名或密码不能为空")
	}
	user := dao.FindUserByUserName(username)
	password = strings.ToUpper(utils.Md5(password + user.Salt))
	if user.Password != password {
		return "密码不正确", errors.New("密码不正确")
	}
	if user.Status == 0 {
		return "您的账号被禁用", errors.New("您的账号被禁用")
	}
	token := addLoginToken(user.Uid)
	return token, nil
}
/**
用户注册 todo email
*/
func Register(username, password string) string {
	if username == "" || password == "" {
		return "用户名或密码不能为空"
	}
	user := dao.FindUserByUserName(username)
	// todo 验证
	if user != (models.User{}) {
		return "用户名已经存在"
	}
	//获取salt
	uuids, err := uuid.NewV4()
	utils.CheckError(err)
	runes := []rune(uuids.String())
	salt := make([]rune, 5)
	for i := 0; i < 5; i++{
		salt[i] = runes[i]
	}
	user = models.User{
		Username: username,
		Password: strings.ToUpper(utils.Md5(password + string(salt))) ,
		Salt: string(salt),
		Status: 1,
		Nickname: "footmap用户",
		Manager: 0,
	}
	dao.AddUser(user)
	//todo 绑定邮箱
	return "注册成功"
}
 /**
给用户下发token
 */
func addLoginToken(uid int) string {
	uuid, err := uuid.NewV4()
	utils.CheckError(err)
	token := models.Token{
		Uid: uid,
		Status: 1,
		Tokens: strings.Replace(uuid.String(),"-","",-1),
		Expired: time.Now().Add(3 * 24 * time.Hour),
	}
	dao.AddToken(token)
	return token.Tokens
}
/**
退出
 */
 func Logout(token string) {
	 dao.UpdateTokenStatus(token, 0)
 }
/**
获取当前用户信息
 */
 func GetUser(token string) models.ResponseUser{
 	var u models.User
 	var respu models.ResponseUser
 	uid := dao.GetUserByToken(token)
	//根据uid查询用户
	 u = dao.GetUserByUid(uid)
	 respu = models.ResponseUser{
	 	Nickname: u.Nickname,
	 	Manager: u.Manager,
	 }
	 return respu
 }
/**
修改用户状态
 */
func UpdateStatus(username string, status int) {
	dao.UpdateUserStatus(username, status)
}
/**
修改用户昵称和密码
 */
 func UpdateUser(nickname, password, token string) {
	 uid := dao.GetUserByToken(token)
	 user := dao.GetUserByUid(uid)
	 password = strings.ToUpper(utils.Md5(password + user.Salt))
	 dao.UpdateUser(nickname, password, uid)
 }