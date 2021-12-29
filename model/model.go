package model

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)
const (
	GetTokenUrl = "https://oapi.dingtalk.com/gettoken"                     //获取token信息接口
	GetListSubIdUrl = "https://oapi.dingtalk.com/topapi/v2/department/listsubid" //获取部门list信息
	GetDepDetailUrl = "https://oapi.dingtalk.com/topapi/v2/department/get"   //获取部门详细信息接口
	GetUserListUrl="https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/queryonjob"    //获取用户list信息
	GetUserDetailUrl="https://oapi.dingtalk.com/topapi/v2/user/get"  //获取部门详细信息


)

type LDAP  interface{
	AddGroupinfo()
	AddUserinfo()
}

var (
	Token             string                                    //接口token信息
	METHOD            = "GET"                               //请求方法
	DepID             =1                                    //默认department ID信息
	ScheduTimeChan    chan bool           //定时器触发器
)

type Tokenstr struct {
	Errcode      int    `json:"errocd"`
	Access_token string `json:"access_token"`
	Errmsg       string `json:"errmsg"`
	Expires_in   int    `json:"expires_in"`
	//info []interface{}
}


func UrlRequest(method string, url string, body *map[string]interface{}) ([]byte,error) {
	C := &http.Client{}
	var bodyinfo io.Reader=nil
	if body != nil {
		//for k, v := range *body {
		//	log.Println(k, v)
		//}
		json_body,err:=json.Marshal(body)
		if err != nil{
			log.Println(err)
			return nil,err
		}else{
			bodyinfo=bytes.NewReader(json_body)
		}
	}

	resq, err := http.NewRequest(method, url, bodyinfo)
	if err != nil {
		log.Error(err)
		return nil,err
	}
	//defer resq.Body.Close()

	res, err := C.Do(resq)
	defer res.Body.Close()
	if err != nil {
		log.Error(err)
	}
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
	}
	return str,nil
}

func GetToken(method string) {
	Url_token:=GetTokenUrl+"?appkey="+ Authconfig.AppKey+"&appsecret=" + Authconfig.AppSecret
	str,_ := UrlRequest(method, Url_token, nil)
	json_info := Tokenstr{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
	}
	//token := json_info.Access_token
	Token = json_info.Access_token
	//return token
}

func Clone(a, b interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(a); err != nil {
		return err
	}
	if err := dec.Decode(b); err != nil {
		return err
	}
	return nil
}


