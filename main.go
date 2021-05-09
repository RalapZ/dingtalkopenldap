package main

import (
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"github.com/RalapZ/dingtalkopenldap/service"
	"strconv"
)

func main() {
	AppKey := "dingicgycuisw7lrqcj8"
	AppSecret := "pUTLevcbvk9EWdJKF7kzj5-Txc_7CHdyrO9eFiObLj7Qb6F3Y7q8YfXnwhtIFTvI"
	method := "GET"
	Url_token := "https://oapi.dingtalk.com/gettoken?appkey=" + AppKey + "&appsecret=" + AppSecret
	service.GetToken(method, Url_token)
	fmt.Println(model.Token)

	//获取所有部门信息https://oapi.dingtalk.com/topapi/v2/department/listsub
	UrlListSub := "https://oapi.dingtalk.com/department/list?access_token=" + model.Token
	sub := service.GetListSub(method, UrlListSub)
	fmt.Println(sub)

	UrlDepDetail := "https://oapi.dingtalk.com/topapi/v2/department/get?access_token=" + model.Token + "&dept_id="
	for _, k := range sub {
		url := UrlDepDetail + strconv.Itoa(k.Id)
		subdetail := service.GetSubDetailInfo(method, url)
		fmt.Printf("%v\n", subdetail)
		//fmt.Printf("%+#v\n",subdetail)
	}

	//GetListSubId
}
