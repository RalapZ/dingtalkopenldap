package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	UserListDetailInfo     =make(map[string]UserDetailInfo)          //user 详细信息
	Userlist   []string
)


type ResponUserInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  struct {
		DataList []string `json:"data_list"`
	} `json:"result"`
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}

type UserDetailInfo struct {
	Active        bool   `json:"active"`
	Admin         bool   `json:"admin"`
	Avatar        string `json:"avatar"`
	Boss          bool   `json:"boss"`
	DeptIdList    []int  `json:"dept_id_list"`
	DeptOrderList []struct {
		DeptId int   `json:"dept_id"`
		Order  int64 `json:"order"`
	} `json:"dept_order_list"`
	Email            string `json:"email"`
	ExclusiveAccount bool   `json:"exclusive_account"`
	HideMobile       bool   `json:"hide_mobile"`
	JobNumber        string `json:"job_number"`
	LeaderInDept     []struct {
		DeptId int  `json:"dept_id"`
		Leader bool `json:"leader"`
	} `json:"leader_in_dept"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	RealAuthed bool   `json:"real_authed"`
	Remark     string `json:"remark"`
	Senior     bool   `json:"senior"`
	StateCode  string `json:"state_code"`
	Telephone  string `json:"telephone"`
	Title      string `json:"title"`
	Unionid    string `json:"unionid"`
	Userid     string `json:"userid"`
	WorkPlace  string `json:"work_place"`
}

type ResponUserDetailinfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  UserDetailInfo `json:"result"`
	RequestId string `json:"request_id"`
}

//获取用户的list
func GetListUserInfoMap(method string,offset int,url string){
	body:= make(map[string]interface{})
	//body["status_list"]=Defaultconfig.StatusList
	body["status_list"]="2,3,5,-1"
	body["offset"]= offset
	atoi, err2 := strconv.Atoi(Defaultconfig.UserOffset)
	if err2 !=nil {
		log.Println(err2)
	}
	body["size"] = atoi
	str := UrlRequest(method, url,&body)
	//body_json := err2
	json_info := ResponUserInfo{}
	err := json.Unmarshal(str, &json_info)
	if err != nil {
		log.Println(err)
	}
	if len(json_info.Result.DataList)!=0{
		Userlist=append(Userlist,json_info.Result.DataList...)
		log.Println(Userlist)
		GetListUserInfoMap(method ,offset+atoi,url)
	}
}

//获取用户详细信息
func GetUserDetailInfo(method string,UserID string, url string) {
	body:= make(map[string]interface{})
	body["userid"]=UserID
	str := UrlRequest(method, url,&body)
	//body_json := err2
	json_info := ResponUserDetailinfo{}
	err := json.Unmarshal(str, &json_info)
	if _,ok:=UserListDetailInfo[UserID];!ok {
		UserListDetailInfo[UserID] = json_info.Result
	}
	//fmt.Println(json_info)
	if err != nil {
		log.Println(err)
	}
}

