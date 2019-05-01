package dao

import (
	"footmap/app/models"
	"footmap/app/utils"
	"time"
)
func AddToken(token models.Token) {
	getConn()
	stmt, err := db.Prepare("insert into token (uid, status, expired, tokens) values (?, ?, ?, ?)")
	utils.CheckError(err)
	_, err = stmt.Exec(token.Uid, token.Status, token.Expired, token.Tokens)
	utils.CheckError(err)
	err = db.Close()
	utils.CheckError(err)
}

func UpdateTokenStatus(token string, status int) {
	getConn()
	stmt, err := db.Prepare("update token set status = ? where tokens = ?")
	utils.CheckError(err)
	_, err = stmt.Exec(status , token)
	utils.CheckError(err)
	err = db.Close()
	utils.CheckError(err)
}

func GetUserByToken(token string) int {
	getConn()
	rows, err := db.Query("select uid from token where tokens = ?", token)
	utils.CheckError(err)
	var uid int
	for rows.Next() {
		err := rows.Scan(&uid)
		utils.CheckError(err)
	}
	err = db.Close()
	utils.CheckError(err)
	return uid
}
func GetToken(token string) models.Token {
	getConn()
	rows, err := db.Query("select uid, status, expired from token where tokens = ?", token)
	utils.CheckError(err)
	var t models.Token
	for rows.Next() {
		var uid, status int
		var expired string
		err := rows.Scan(&uid, &status, &expired)
		expiredtime, err := time.Parse("2006-01-02", expired)
		utils.CheckError(err)
		t = models.Token{
			Uid: uid,
			Expired: expiredtime,
			Status: status,
			Tokens: token,
		}
	}
	err = db.Close()
	return t
}