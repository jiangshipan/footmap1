package dao

import (
	"database/sql"
	"footmap/app/models"
	"footmap/app/utils"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB
// init在main加载的时候就会执行，所以用单独的函数
func getConn() {
	var err error
	db, err = sql.Open("mysql", "root:qqq110@/footmap")
	utils.CheckError(err)
}
//根据用户名查询用户
func FindUserByUserName(username string) models.User {
	//得到数据库连接
	getConn()
	var user models.User
	rows, err := db.Query("select uid,password,nickname,salt,status,manager from user where username = ?", username)
	utils.CheckError(err)
	for rows.Next() {
		var uid, status, manager int
		var password, nickname, salt string
		err := rows.Scan(&uid, &password, &nickname, &salt, &status, &manager)
		utils.CheckError(err)
		user = models.User{
			Uid: uid,
			Username: username,
			Password: password,
			Nickname: nickname,
			Salt: salt,
			Status: status,
			Manager: manager,
		}
	}
	err = db.Close()
	utils.CheckError(err)
	return user
}
//添加用户
func AddUser(user models.User) {
	getConn()
	stmt, err := db.Prepare("insert into user (username,password,nickname,salt,status,manager) values (?,?,?,?,?,?)")
	utils.CheckError(err)
	_, err = stmt.Exec(user.Username, user.Password, user.Nickname, user.Salt, user.Status, user.Manager)
	utils.CheckError(err)
}

func GetUserByUid(uid int) models.User {
	getConn()
	rows, err := db.Query("select nickname,salt,manager from user where uid = ?", uid)
	utils.CheckError(err)
	var user models.User
	for rows.Next() {
		var manager int
		var nickname, salt string
		err := rows.Scan(&nickname, &salt, &manager)
		utils.CheckError(err)
		user = models.User{
			Nickname: nickname,
			Salt: salt,
			Manager: manager,
		}
	}
	err = db.Close()
	utils.CheckError(err)
	return user
}

func UpdateUserStatus(username string, status int) {
	getConn()
	stmt, err := db.Prepare("update user set status = ? where username = ?")
	utils.CheckError(err)
	_, err = stmt.Exec(status, username)
	utils.CheckError(err)
	err = db.Close()
	utils.CheckError(err)
}

func UpdateUser(nickname, password string, uid int) {
	getConn()
	stmt, err := db.Prepare("update user set nickname = ?, password = ? where uid = ?")
	utils.CheckError(err)
	_, err = stmt.Exec(nickname, password, uid)
	utils.CheckError(err)
	err = db.Close()
	utils.CheckError(err)
}