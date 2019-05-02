package dao

import (
	"footmap/app/models"
	"footmap/app/utils"
	"time"
)

func AddRecord(record models.Record) {
	getConn()
	stmt, err := db.Prepare("insert into record (uid,startTime,targetTime,remarks,targetPlace,title,status) values (?,?,?,?,?,?,?)")
	utils.CheckError(err)
	_, err = stmt.Exec(record.Uid, record.StartTime, record.TargetTime, record.Remarks, record.TargetPlace, record.Title, record.Status)
	utils.CheckError(err)
	err = db.Close()
	utils.CheckError(err)
}

func GetRecord() []models.Record {
	getConn()
	rows, err := db.Query("select uid, startTime,targetTime,remarks,targetPlace,title from record where status=1 and targetTime > now() order by targetTime asc limit 5")
	utils.CheckError(err)

	var records  = make([]models.Record, 0)
	for rows.Next() {
		var uid int
		var remarks,targetPlace,title,startTime, targetTime string
		err := rows.Scan(&uid, &startTime, &startTime, &remarks, &targetPlace, &title)
		utils.CheckError(err)
		stime, err := time.Parse("2006-01-02 15:04:05", startTime)
		ttime, err := time.Parse("2006-01-02 15:04:05", targetTime)

		record := models.Record{
			Uid: uid,
			Remarks: remarks,
			TargetPlace: targetPlace,
			Title: title,
			StartTime: stime,
			TargetTime: ttime,
		}
		records = append(records, record)
	}
	err = db.Close()
	utils.CheckError(err)
	return records
}