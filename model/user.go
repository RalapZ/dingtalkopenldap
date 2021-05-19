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
	Active        bool   `json:"active"`   //是否激活了钉钉： true：已激活false：未激活
	Admin         bool   `json:"admin"`   //是否为企业的管理员： true：是false：不是
	Avatar        string `json:"avatar"`  //头像
	Boss          bool   `json:"boss"`    //是否为企业的老板： true：是false：不是
	DeptIdList    []int  `json:"dept_id_list"`    //所属部门ID列表。
	DeptOrderList []struct {           //员工在对应的部门中的排序。
		DeptId int   `json:"dept_id"`    //部门ID。
		Order  int64 `json:"order"`      //员工在部门中的排序。
	} `json:"dept_order_list"`
	Email            string `json:"email"`     //员工邮箱。
	ExclusiveAccount bool   `json:"exclusive_account"` //是否专属帐号。
	HideMobile       bool   `json:"hide_mobile"`    //是否号码隐藏： true：隐藏false：不隐藏
	JobNumber        string `json:"job_number"`     //员工工号
	LeaderInDept     []struct {                  //员工在对应的部门中是否领导。
		DeptId int  `json:"dept_id"`
		Leader bool `json:"leader"`
	} `json:"leader_in_dept"`
	Mobile     string `json:"mobile"`     //手机号码。
	Name       string `json:"name"`       //员工名称。
	RealAuthed bool   `json:"real_authed"`    //是否完成了实名认证： true：已认证false：未认证
	Remark     string `json:"remark"`    //备注。
	Senior     bool   `json:"senior"`    //是否为企业的高管： true：是false：不是
	StateCode  string `json:"state_code"`   //国际电话区号。
	Telephone  string `json:"telephone"`    //分机号
	Title      string `json:"title"`    //职位
	Unionid    string `json:"unionid"`   //员工在当前开发者企业账号范围内的唯一标识
	Userid     string `json:"userid"`    //员工的userid
	WorkPlace  string `json:"work_place"`    //办公地点。
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


func UserCompareInfo(srcuser UserDetailInfo,dstuser UserDetailInfo) bool{
	ComRe,_:= OrderIntArrayCompare(srcuser.DeptIdList, dstuser.DeptIdList)
	if ComRe==false&&srcuser.Name==dstuser.Name && srcuser.Telephone==dstuser.Telephone && srcuser.Email==dstuser.Email{
		return false
	}
	return true
}


func OrderIntArrayCompare(srcarr []int,dstarr []int)(bool,error){
	//if reflect.TypeOf(srcarr).Kind()!=reflect.Slice || reflect.TypeOf(dstarr).Kind()!=reflect.Slice{
	//	return false,errors.New("one of args is not slice")
	//}
	//if reflect.TypeOf(srcarr[0]).Elem()!= reflect.TypeOf(dstarr[0]).Elem(){
	//	return false,errors.New("slice type is different")
	//}
	if len(srcarr)!=len(dstarr){
		return true,nil
	}
	for i:=0;i<len(srcarr);i++{
		if srcarr[i]!=dstarr[i]{
			return true,nil
		}
	}
	return false,nil
}