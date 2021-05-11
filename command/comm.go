package command

import (
	"github.com/RalapZ/dingtalkopenldap/model"
	"log"
	"time"
)

func InitLdap(){
	//fmt.Println(model.AuthConf, model.Listenconfig, model.Authconfig, model.Ldapconfig, model.DBconfig)
	//model.InitLdapConnection()
	//AppKey := "dingicgycuisw7lrqcj8"
	//AppSecret := "pUTLevcbvk9EWdJKF7kzj5-Txc_7CHdyrO9eFiObLj7Qb6F3Y7q8YfXnwhtIFTvI"
	//Url_token := "https://oapi.dingtalk.com/gettoken?appkey=" + AppKey + "&appsecret=" + AppSecret
	//Url_token:=model.GetTokenUrl+"?appkey="+ model.Authconfig.AppKey+"&appsecret=" + model.Authconfig.AppSecret
	model.GetToken("GET")
	log.Println(model.Token)
	model.GetListUserInfoMap("POST",0,model.GetUserListUrl+"?access_token="+"d17128d9cd2a3f188e32abbdd7ace8a0")
	UrlDepSubId := "https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=" + model.Token
	//model.GetListSubId("POST",1 , UrlDepSubId)
	model.InitListSubId("POST",1 , UrlDepSubId)
	//GetListSubId
	time.Sleep(1000*time.Microsecond)
	log.Println(model.DepListId)
	//for k,v:=range model.DepListDetailInfo{
	//	log.Println(k,v)
	//}
	for k,v:=range model.UserListDetailInfo{
		log.Println(k,v)
	}

}
