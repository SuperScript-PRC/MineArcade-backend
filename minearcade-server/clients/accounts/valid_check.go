package accounts

import "strings"

func AccountNameValid(accountName string) (bool, string) {
	if len(accountName) < 4 {
		return false, "账号名太短"
	} else if len(accountName) > 16 {
		return false, "账号名太长"
	} else if strings.HasPrefix(accountName, "__") {
		return false, "无效账号名"
	}
	return true, ""
}

func NickNameValid(nickName string) (bool, string) {
	if len(nickName) < 1 {
		return false, "昵称太短"
	} else if len(nickName) > 15 {
		return false, "昵称太长"
	} else if strings.Contains(nickName, " ") {
		return false, "无效昵称名"
	}
	return true, ""
}
