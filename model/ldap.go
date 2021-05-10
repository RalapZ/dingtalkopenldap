package model

import (
	"fmt"
	log  "github.com/sirupsen/logrus"

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
	//sql := ldap.NewAddRequest("ou=it,dc=asinking,dc=com")
	//sql.Attribute("objectClass", []string{"organizationalUnit"})
	//er := LDAPservice.Conn.Add(sql)
	//if er!=nil{
	//	fmt.Println(er)
	//}
}



//
//func ErrHanding(test func()) func(){
//	return func(){
//
//	}
//}


func (ldapservice *LDAPService)AddGroupinfo(groupinfo DepDetailInfo) {
	//sql := ldap.NewAddRequest("ou=开发部3,ou=it,dc=asinking,dc=com")
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

func (ldapservice *LDAPService)AddGroupinfotest(){
	sql := ldap.NewAddRequest("o=AsinKing,dc=asinking,dc=com")
	sql.Attribute("ou", []string{"AsinKing"})
	sql.Attribute("objectClass", []string{"organization","top"})
	er := LDAPservice.Conn.Add(sql)
	if er!=nil{
		log.Println(er)
	}
}

func (ldapservice *LDAPService)AddUserinfo(){
	sql := ldap.NewAddRequest("uid=test,ou=people,ou=it,dc=asinking,dc=com")
	//sql := ldap.NewAddRequest("ou=供应链组,ou=产品部,ou=Staff,ou=Groups,o=AsinKing,dc=asinking,dc=com")

	sql.Attribute("uidNumber", []string{"1010"})
	sql.Attribute("gidNumber", []string{"1003"})
	sql.Attribute("userPassword", []string{"123456"})
	sql.Attribute("homeDirectory", []string{"/home/wujq"})
	sql.Attribute("cn", []string{"test"})
	sql.Attribute("uid", []string{"test"})
	sql.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	//sql.Attribute("objectClass", []string{"inetOrgPerson","organizationalPerson","person","top"})
	er := LDAPservice.Conn.Add(sql)
	if er!=nil{
		fmt.Println(er)
	}
}







