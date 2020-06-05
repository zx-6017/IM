package helpers

import "strings"

func GetUidByIdentify(identify string) string{

	uid_arr := strings.Split(identify,"_")
	uid := ""
	if len(uid_arr) == 3 {
		uid = uid_arr[2]
	}else {
		uid = uid_arr[1]
	}
	return uid
}
