package model

import "testing"

func TestGetUserDetailInfo(t *testing.T) {
	//url1:=GetUserDetailUrl+"?access_token="+"9c9e19ee012e3eaeb27e9717d2f116ef"
	//	//"https://oapi.dingtalk.com/topapi/v2/user/get?access_token=9c9e19ee012e3eaeb27e9717d2f116ef"
	//GetUserDetailInfo("POST","012609262508120583", url1)
	//url2:=GetUserListUrl+"?access_token="+"d17128d9cd2a3f188e32abbdd7ace8a0"
	//GetListUserInfoMap("POST",0,url2)
	GetListUserInfoMap("POST",0,GetUserListUrl+"?access_token="+"576916ba412836bc8d0c8534055c0aa0")
}

func TestCheckUserInfo(t *testing.T) {
	test:="https://oapi.dingtalk.com/topapi/smartwork/hrm/employee/queryonjob?access_token="+"ec59a699580e353eaac38318e7ca08a6"
	//setConf()
	for {
		CheckUserInfo("POST",0,test)
	}

}