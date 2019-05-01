package controller

import (
	"encoding/json"
	"footmap/app/dao"
	"footmap/app/models"
	"footmap/app/utils"
	"net/http"
	"strings"
	"time"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	if split[2] == "login" || split[2] == "reg" {
		//对login和reg 不验证身份
		UserHandler(w, r)
	} else {
		var resp models.MyResponse
		cookie, err := r.Cookie("auth_token")
		//没有auth_token
		if err == http.ErrNoCookie {
			resp = models.MyResponse{
				Code: 1,
				Msg: "请登录",
				Data: struct{}{},
			}
			bytes, err := json.Marshal(resp)
			utils.CheckError(err)
			w.Write(bytes)
		} else {
			//根据cookie查询过期时间
			token := dao.GetToken(cookie.Value)
			if token != (models.Token{}) || token.Status != 0 || token.Expired.Before(time.Now()) {
				//验证成功
				//选择调用那个处理器 split[1]
				switch split[1] {
				case "user": UserHandler(w, r)
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

				//			UserHandler(w, r)
			} else {
				resp = models.MyResponse{
					Code: 1,
					Msg: "请登录",
					Data: struct{}{},
				}
				bytes, err := json.Marshal(resp)
				utils.CheckError(err)
				w.Write(bytes)
			}
	}

	}

}