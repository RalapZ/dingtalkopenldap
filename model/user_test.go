package model

import "testing"

func TestGetUserDetailInfo(t *testing.T) {
	//url1:=GetUserDetailUrl+"?access_token="+"9c9e19ee012e3eaeb27e9717d2f116ef"
	//	//"https://oapi.dingtalk.com/topapi/v2/user/get?access_token=9c9e19ee012e3eaeb27e9717d2f116ef"
	//GetUserDetailInfo("POST","012609262508120583", url1)
	//url2:=GetUserListUrl+"?access_token="+"d17128d9cd2a3f188e32abbdd7ace8a0"
	//GetListUserInfoMap("POST",0,url2)
	GetListUserInfoMap("POST",0,GetUserListUrl+"?access_token="+"d17128d9cd2a3f188e32abbdd7ace8a0")
}