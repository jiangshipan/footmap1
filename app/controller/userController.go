package controller

import (
	"encoding/json"
	"footmap/app/models"
	"footmap/app/services"
	"footmap/app/utils"
	"net/http"
	"strconv"
	"strings"
)

//user 处理器
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": handleGet(w, r)
	case "POST": handlePost(w, r)
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
func handleGet(w http.ResponseWriter, r *http.Request) {
	//响应
	var resp models.MyResponse
	// 获取/user/xxx
	url := r.URL.Path
	//若/在前面，则数组大小为3，否则为2  若url为user/login 则切割后数组长度为2
	str := strings.Split(url, "/")
	//xxx
	keywords := str[2]

	switch keywords {
	case "login":
		username := utils.GetParam(r.RequestURI, "username")
		password := utils.GetParam(r.RequestURI, "password")
		message, err := services.Login(username, password)
		if err == nil {
			//登陆成功
			cookie := http.Cookie{
				Name: "auth_token",
				Value: message,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			resp = models.MyResponse{
				Code: 0,
				Msg: "登陆成功",
				Data: struct{}{},
			}
		} else {
			utils.CheckError(err)
			resp = models.MyResponse{
				Code: 1,
				Msg: message,
				Data: struct{}{},
			}
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)

	case "logout":
		cookie, err := r.Cookie("auth_token")
		utils.CheckError(err)
		services.Logout(cookie.Value)
		resp = models.MyResponse{
			Code: 1,
			Msg: "退出成功",
			Data: struct{}{},
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)

	case "getUser":
		cookie, err := r.Cookie("auth_token")
		utils.CheckError(err)
		user := services.GetUser(cookie.Value)
		resp = models.MyResponse{
			Code: 0,
			Msg: "获取成功",
			Data: user,
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)

	case "updateStatus":
		username := utils.GetParam(r.RequestURI, "username")
		status := utils.GetParam(r.RequestURI, "status")
		i, err := strconv.Atoi(status)
		utils.CheckError(err)
		services.UpdateStatus(username, i)
		resp := models.MyResponse{
			Code: 0,
			Msg: "修改成功",
			Data: struct{}{},
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)
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
	//fmt.Println(url)
}
//post请求处理器
func handlePost(w http.ResponseWriter, r *http.Request) {
	var resp models.MyResponse
	// 获取/user/xxx
	url := r.URL.Path
	//若/在前面，则数组大小为3，否则为2  若url为user/login 则切割后数组长度为2
	str := strings.Split(url, "/")
	//xxx
	keywords := str[2]
	switch keywords {
	case "reg":
		jsonMap := utils.GetOneField(r)
		username := jsonMap["username"]
		password := jsonMap["password"]//interface{} -> 具体类型 password.(string)
		//email := utils.GetParam(r.RequestURI, "email")
		message := services.Register(username.(string), password.(string))
		if message == "注册成功" {
			resp = models.MyResponse {
				Code: 0,
				Msg: message,
				Data: struct{}{},
			}
		} else {
			resp = models.MyResponse{
				Code: 1,
				Msg: message,
				Data: struct{}{},
			}
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)

	case "updateUser":
		jsonMap := utils.GetOneField(r)
		nickname := jsonMap["nickname"]
		password := jsonMap["password"]
		cookie, err := r.Cookie("auth_token")
		utils.CheckError(err)
		if nickname == "" || password == "" {
			resp := models.MyResponse{
				Code: 1,
				Msg: "密码不能为空",
				Data: struct{}{},
			}
			bytes, err := json.Marshal(resp)
			utils.CheckError(err)
			w.Write(bytes)
		}
		services.UpdateUser(nickname.(string), password.(string), cookie.Value)
		resp := models.MyResponse{
			Code: 0,
			Msg: "修改成功",
			Data: struct{}{},
		}
		bytes, err := json.Marshal(resp)
		utils.CheckError(err)
		w.Write(bytes)
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