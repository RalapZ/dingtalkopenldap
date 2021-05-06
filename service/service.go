package service

import (
	"encoding/json"
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"io/ioutil"
	"net/http"
)

func UrlRequest(method string, url string) []byte {
	C := &http.Client{}
	resq, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	//defer resq.Body.Close()

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

func GetToken(method string, url string) string {
	str := UrlRequest(method, url)
	json_info := model.Tokenstr{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		fmt.Println(err)
	}
	token := json_info.Access_token
	model.Token = json_info.Access_token
	return token
}

func GetListSub(method string, url string) []model.SubInfo {
	str := UrlRequest(method, url)
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
	str := UrlRequest(method, url)
	json_info := model.DepDetailInfo{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		fmt.Println(err)
	}
	//token:=json_info.Access_token
	//DepartmentInfo:=
	return json_info
}
