package services

import (
	"footmap/app/dao"
	"footmap/app/models"
	"footmap/app/utils"
	"time"
)

func Addrecord(jsonMap map[string]interface{}, token string) string {
	startTime := jsonMap["startTime"]
	targetTime := jsonMap["targetTime"]
	remarks := jsonMap["remarks"]
	targetPlace := jsonMap["targetPlace"]
	title := jsonMap["title"]
	if (startTime == nil || targetTime == nil || targetPlace == nil) {
		return "参数不完整"
	}
	sTime := startTime.(string)
	tTime := targetTime.(string)
	stime, err := time.Parse("2006-01-02 15:04:05", sTime)
	utils.CheckError(err)
	ttime, err := time.Parse("2006-01-02 15:04:05", tTime)
	utils.CheckError(err)
	uid := dao.GetUserByToken(token)
	r := models.Record{
		Uid: uid,
		StartTime: stime,
		TargetTime: ttime,
		Remarks: remarks.(string),
		TargetPlace: targetPlace.(string),
		Title: title.(string),
		Status: 1,
	}
	dao.AddRecord(r)
	return "新增成功"

}