package service

import (
	"encoding/json"
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"io/ioutil"
	"net/http"
)

var DDtoken string


func UrlRequest(method string, url string,body *map[interface{}]interface{}) []byte {
	C := &http.Client{}
	resq, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	//defer resq.Body.Close()
	if body == nil{
		fmt.Println("body is null")
	}else{
		for k,v := range *body{
			fmt.Println(k,v)
		}
		//resq.Header.Add()
	}

	res, err := C.Do(resq)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return str
}

func GetToken(method string, url string) {
	str := UrlRequest(method, url,nil)
	json_info := model.Tokenstr{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		fmt.Println(err)
	}
	//token := json_info.Access_token
	model.Token = json_info.Access_token
	//return token
}

func GetListSubId(method string,DepID int, url string) {
	body:= map[interface{}]interface{}
	body["dept_id"]=DepID
	str := UrlRequest(method, url,&body)
	json_info := model.ResponseDepListSubId{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

func GetListSub(method string, url string) []model.SubInfo {
	str := UrlRequest(method, url,nil)
	json_info := model.ListSubInfo{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		fmt.Println(err)
	}
	//token:=json_info.Access_token
	//DepartmentInfo:=
	return json_info.Department
}

func GetSubDetailInfo(method string, url string) model.DepDetailInfo {
	str := UrlRequest(method, url,nil)
	json_info := model.DepDetailInfo{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		fmt.Println(err)
	}
	//token:=json_info.Access_token
	//DepartmentInfo:=
	return json_info
}
