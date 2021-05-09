package model

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"testing"
)


var test LDAPService

func TestInitLdapConnection(t *testing.T) {
	Ldapconfig=LdapConfig{
		Addr :"192.168.137.201:389",
		BindUserName:"cn=manager,dc=asinking,dc=com",
		BindPassword :"1qaz@WSX",
		SearchDN :"dc=asinking,dc=com",
	}
	InitLdapConnection()
	searchRequest := ldap.NewSearchRequest(Ldapconfig.SearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", "op01"),
		[]string{"dn"},
		nil,)
	sr, err := LDAPservice.Conn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}
	sr.Print()

}
//
