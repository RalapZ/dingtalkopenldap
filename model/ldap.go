package model

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
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





