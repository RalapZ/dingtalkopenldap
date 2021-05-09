package model

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
const (
	GetTokenUrl = "https://oapi.dingtalk.com/gettoken"                     //获取token信息接口
	GetDepDetailUrl = "https://oapi.dingtalk.com/topapi/v2/department/get"   //获取部门详细信息接口
)

var (
	Token             string                                    //接口token信息
	METHOD                = "GET"                               //请求方法

	DepID             int =1                                    //默认department ID信息
)

type Tokenstr struct {
	Errcode      int    `json:"errocd"`
	Access_token string `json:"access_token"`
	Errmsg       string `json:"errmsg"`
	Expires_in   int    `json:"expires_in"`
	//info []interface{}
}






func UrlRequest(method string, url string, body *map[string]interface{}) []byte {
	C := &http.Client{}
	//bodyinfo:=nil
	var bodyinfo io.Reader=nil
	if body != nil {
		for k, v := range *body {
			log.Println(k, v)
		}
		json_body,err:=json.Marshal(body)
		if err != nil{
			log.Println(err)
		}else{
			bodyinfo=bytes.NewReader(json_body)
		}
	}
	resq, err := http.NewRequest(method, url, bodyinfo)
	if err != nil {
		log.Println(err)
	}
	//defer resq.Body.Close()

	res, err := C.Do(resq)
	defer res.Body.Close()
	if err != nil {
		log.Println(err)
	}
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	return str
}

func GetToken(method string, url string) {
	str := UrlRequest(method, url, nil)
	json_info := Tokenstr{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
	}
	//token := json_info.Access_token
	Token = json_info.Access_token
	//return token
}

