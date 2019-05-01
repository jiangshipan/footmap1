package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
获取json 返回map
 */
func GetOneField(r *http.Request) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	jsonStr := bytes.NewBuffer(body).String()
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	CheckError(err)
	return jsonMap
}