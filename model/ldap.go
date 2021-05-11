package model

import (
	"fmt"
	log  "github.com/sirupsen/logrus"
	"strconv"

	//"github.com/go-ldap/ldap/v3"
	"gopkg.in/ldap.v2"
)

//var LdapConn *ldap.Conn

var LDAPservice LDAPService

type LDAPService struct {
	Conn     *ldap.Conn
	Config     LdapConfig
}

func InitLdapConnection(){
	conn, err := ldap.Dial("tcp", Ldapconfig.Addr)

	if err!=nil{
		panic(err)
	}
	LDAPservice=LDAPService{conn,Ldapconfig}
	err = LDAPservice.Conn.Bind(Ldapconfig.BindUserName, Ldapconfig.BindPassword)
	if err !=nil{
		fmt.Println(err)
		panic(err)
	}
}


func (ldapservice *LDAPService)AddGroupinfo(groupinfo DepDetailInfo) {
	var grouptempinfo string
	for _,k:=range StackDepmentinfo{
		grouptempinfo="ou="+k+","+grouptempinfo
	}
	fmt.Println(grouptempinfo)
	sql := ldap.NewAddRequest(grouptempinfo+Ldapconfig.SearchDN)

	sql.Attribute("ou", []string{StackDepmentinfo[len(StackDepmentinfo)-1]})
	sql.Attribute("objectClass", []string{"organizationalUnit","top"})
	er := LDAPservice.Conn.Add(sql)
	if er!=nil{
		log.Println(er)
	}
	log.Println()
}


func (ldapservice *LDAPService)AddUserinfo(userid int){
	//var usertempinfo string
	var userldappath string
	for _,k:=range DepListDetailInfo[userid].LdapDepPath{
		userldappath="ou="+k+","+userldappath
	}
	usertempinfo:=ldap.NewAddRequest("uid="+UserListDetailInfo[strconv.Itoa(userid)].Name+","+userldappath+Ldapconfig.SearchDN)
	//sql := ldap.NewAddRequest("uid=test,ou=people,ou=it,dc=asinking,dc=com")
	//sql := ldap.NewAddRequest("ou=供应链组,ou=产品部,ou=Staff,ou=Groups,o=AsinKing,dc=asinking,dc=com")
	usertempinfo.Attribute("uidNumber", []string{UserListDetailInfo[strconv.Itoa(userid)].Userid})
	usertempinfo.Attribute("gidNumber", []string{strconv.Itoa(UserListDetailInfo[strconv.Itoa(userid)].DeptOrderList[0].DeptId)})
	usertempinfo.Attribute("userPassword", []string{"123456"})
	usertempinfo.Attribute("homeDirectory", []string{"/home/"+UserListDetailInfo[strconv.Itoa(userid)].Name})
	usertempinfo.Attribute("cn", []string{"test"})
	usertempinfo.Attribute("uid", []string{UserListDetailInfo[strconv.Itoa(userid)].Userid})
	usertempinfo.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	//sql.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson","person","top"})
	er := LDAPservice.Conn.Add(usertempinfo)
	if er!=nil{
		fmt.Println(er)
	}
}







