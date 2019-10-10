package controllers

import (
	"IM/models"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// 文件操作练习
func GetLogInfo(request *gin.Context){
	log_file := "/Users/zx/Documents/my_setting/multihello.log"

	err, infos:= readLine(log_file)
	fmt.Println(err)
	for _,value := range infos {
		fmt.Println(value)
	}



}
type add_info struct {
	Receive_uid string
	Send_uid string
	From string
}
func readLine(filePath string) (error,[]string) {
	f,err := os.Open(filePath)
	defer f.Close()

	if err != nil{
		return err,nil
	}
	var s_r_info add_info
	br := bufio.NewReader(f)

	var res []string
	for {
		s,_,c := br.ReadLine()
		if c == io.EOF{
			break
		}

		str := string(s)
		str = strings.Trim(str,"[\r\n")
		arr := strings.Split(str," ")
		time := arr[0]
		add_info,err := url.ParseQuery(strings.Trim(arr[1],"info="))
		if err != nil{
			fmt.Println(err)
			return err,nil
		}
		for key,_ := range add_info{
			err := json.Unmarshal([]byte(key),&s_r_info)
			if err != nil{
				fmt.Println("json error ",err)
				return err,nil
			}
		}
		receive_uids_arr := strings.Split(s_r_info.Receive_uid,",")



		//查询用户信息
		user_model := models.UserInfo{}
		user_infos := user_model.GetUserByIds(receive_uids_arr)
		if len(user_infos)==0{
			continue
		}
		for _,user_info := range user_infos{
			res = append(res, time+" "+user_info.Name+" "+strconv.Itoa(int(user_info.Id))+" "+user_info.Sex+" "+user_info.Age+" "+user_info.Timezone)
		}

	}
	return nil,res
}
