package command

import (
	"fmt"
	"github.com/RalapZ/dingtalkopenldap/model"
	"log"
	"time"
)

func Init(){
	model.InitConfig()
	fmt.Println(model.AuthConf, model.Listenconfig, model.Authconfig, model.Ldapconfig, model.DBconfig)
	model.InitLdapConnection()
	//AppKey := "dingicgycuisw7lrqcj8"
	//AppSecret := "pUTLevcbvk9EWdJKF7kzj5-Txc_7CHdyrO9eFiObLj7Qb6F3Y7q8YfXnwhtIFTvI"

	//Url_token := "https://oapi.dingtalk.com/gettoken?appkey=" + AppKey + "&appsecret=" + AppSecret
	Url_token:=model.GetTokenUrl+"?appkey="+ model.Authconfig.AppKey+"&appsecret=" + model.Authconfig.AppSecret
	model.GetToken("GET", Url_token)
	log.Println(model.Token)

	//获取所有部门信息https://oapi.dingtalk.com/topapi/v2/department/listsub
	//UrlListSub := "https://oapi.dingtalk.com/department/list?access_token=" + model.Token
	//sub := model.GetListSub(method, UrlListSub)
	//log.Println(sub)
	//
	//UrlDepDetail := "https://oapi.dingtalk.com/topapi/v2/department/get?access_token=" + model.Token + "&dept_id="
	//for _, k := range sub {
	//	url := UrlDepDetail + strconv.Itoa(k.Id)
	//	subdetail := model.GetSubDetailInfo("POST", url,)
	//	log.Printf("%v\n", subdetail)
	//	//fmt.Printf("%+#v\n",subdetail)
	//}
	UrlDepSubId := "https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=" + model.Token
	//model.GetListSubId("POST",1 , UrlDepSubId)
	model.InitListSubId("POST",1 , UrlDepSubId)
	//GetListSubId
	time.Sleep(1000*time.Microsecond)
	log.Println(model.DepListId)
	for k,v:=range model.DepListDetailInfo{
		fmt.Println(k,v)
	}
}
