package main

import (
	"github.com/RalapZ/dingtalkopenldap/command"
	"github.com/RalapZ/dingtalkopenldap/model"
)


func init(){
	model.InitConfig()
	model.InitLdapConnection()
	//model.GetToken("GET")

	//log.
}

func main() {
	command.InitLdap()
	//http.ListenAndServe("127.0.0.1:9090",nil)
	//fmt.Println("test",model.Userlist)
}