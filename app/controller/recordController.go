package controller

import (
	"encoding/json"
	"footmap/app/models"
	"footmap/app/services"
	"footmap/app/utils"
	"net/http"
	"strings"
)

//record 处理器
func RecordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": recordHandleGet(w, r)
	case "POST": recordHandlePost(w, r)
	default:
		resp := models.MyResponse{
			Code: 1,
			Msg: "暂不支持",
			Data: struct{}{},
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)
	}
}

//get请求处理器
func recordHandleGet(w http.ResponseWriter, r *http.Request) {
	var resp models.MyResponse
	url := r.URL.Path
	str := strings.Split(url, "/")
	keywords := str[2]

	switch keywords {
	default:
		resp = models.MyResponse{
			Code: 1,
			Msg: "暂不支持",
			Data: struct{}{},
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)
	}
}
//post请求处理器
func recordHandlePost(w http.ResponseWriter, r *http.Request) {
	var resp models.MyResponse
	url := r.URL.Path
	str := strings.Split(url, "/")
	keywords := str[2]
	switch keywords {
	case "add":
		jsonMap := utils.GetOneField(r)
		cookie, err := r.Cookie("auth_token")
		utils.CheckError(err)
		result := services.Addrecord(jsonMap, cookie.Value)
		if result == "新增成功" {
			resp = models.MyResponse{
				Code: 0,
				Msg: "新增成功",
				Data: struct{}{},
			}
			bytes, err := json.Marshal(resp)
			utils.CheckError(err)
			w.Write(bytes)
		}
	default:
		resp := models.MyResponse{
			Code: 1,
			Msg: "暂不支持",
			Data: struct{}{},
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)
	}
}