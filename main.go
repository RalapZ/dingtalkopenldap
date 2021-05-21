package main

import (
	"github.com/RalapZ/dingtalkopenldap/command"
	"github.com/RalapZ/dingtalkopenldap/model"
	"github.com/RalapZ/dingtalkopenldap/tools"
	"net/http"
)


func init(){
	tools.InitLog("./log/", "openldap", "gbk")
	model.InitConfig()
	model.InitLdapConnection()
	//logrus.SetFormatter()

	//model.GetToken("GET")

	//log.
}

func main() {
	command.Init()
	//log.Info("测试中文")
	http.ListenAndServe("127.0.0.1:9090",nil)
	//fmt.Println("test",model.Userlist)
}