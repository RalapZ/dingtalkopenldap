package main

import (
	"github.com/RalapZ/dingtalkopenldap/command"
	"github.com/RalapZ/dingtalkopenldap/model"
	"github.com/RalapZ/dingtalkopenldap/tools"
	"net/http"
)


func init(){
	tools.InitLog("./log/", "openldap", "gbk")
	command.ChangeUserInfoFunc()
	model.InitConfig()
	model.InitLdapConnection()

}

func main() {
	//fmt.Println("Test")
	command.Start()
	//fmt.Println("Test")
	//log.Info("测试中文")
	http.ListenAndServe("127.0.0.1:9090",nil)
	//fmt.Println("test",model.Userlist)
}