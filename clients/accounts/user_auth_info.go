package accounts

import (
	"MineArcade-backend/protocol"
)

type UserAuthInfo struct {
	AccountName string
	PasswordMD5 string
	Nickname    string
	UUIDStr     string
}

func (u *UserAuthInfo) Marshal(w *protocol.Writer) {
	w.StringUTF(u.AccountName)
	w.StringUTF(u.PasswordMD5)
	w.StringUTF(u.Nickname)
	w.StringUTF(u.UUIDStr)
}

func (u *UserAuthInfo) Unmarshal(r *protocol.Reader) {
	r.StringUTF(&u.AccountName)
	r.StringUTF(&u.PasswordMD5)
	r.StringUTF(&u.Nickname)
	r.StringUTF(&u.UUIDStr)
}

func IsPasswordCorrect(username string, passwordMD5 string) (bool, string) {
	db = OpenAccountDB()
	raw_data, err := db.Get([]byte(username), nil)
	if err != nil {
		return false, "用户名不存在"
	}
	reader := protocol.NewReader(raw_data)
	user_auth_info := &UserAuthInfo{}
	user_auth_info.Unmarshal(&reader)
	if user_auth_info.PasswordMD5 != passwordMD5 {
		return false, "密码错误"
	} else {
		return true, ""
	}
}

func GetUserAuthInfo(username string) (*UserAuthInfo, bool) {
	db = OpenAccountDB()
	raw_data, err := db.Get([]byte(username), nil)
	if err != nil {
		return nil, false
	}
	reader := protocol.NewReader(raw_data)
	user_auth_info := &UserAuthInfo{}
	user_auth_info.Unmarshal(&reader)
	return user_auth_info, true
}
