package accounts

import (
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/protocol"
	"strings"

	"github.com/df-mc/goleveldb/leveldb"
	"github.com/pterm/pterm"
)

var db *leveldb.DB

func OpenAccountDB() *leveldb.DB {
	if db == nil {
		// pterm.Info.Println("正在读取账号数据库")
		ldb, err := leveldb.OpenFile(defines.ACCOUNT_DB_PATH, nil)
		if err != nil {
			panic(err)
		}
		db = ldb
	}
	if _, err := db.Get([]byte("admin"), nil); err != nil {
		ud := &UserAuthInfo{AccountName: "admin", Nickname: "admin", PasswordMD5: "1qRG9tE+TwwcupTrbKLT9AAA", UIDStr: "0"}
		w := protocol.NewWriter()
		ud.Marshal(&w)
		db.Put([]byte(ud.AccountName), w.GetFullBytes(), nil)
		pterm.Info.Println("管理员账号未初始化, 已进行初始化:", err)
	}
	if _, err := db.Get([]byte("__UIDTotal"), nil); err != nil {
		b := protocol.NewWriter()
		b.Double(0)
		db.Put([]byte("__UIDTotal"), b.GetFullBytes(), nil)
	}
	return db
}

func GetUserAuthInfo(username string) (*UserAuthInfo, bool) {
	db = OpenAccountDB()
	if strings.HasPrefix(username, "__") {
		return nil, false
	}
	raw_data, err := db.Get([]byte(username), nil)
	if err != nil {
		return nil, false
	}
	reader := protocol.NewReader(raw_data)
	user_auth_info := &UserAuthInfo{}
	user_auth_info.Unmarshal(&reader)
	return user_auth_info, true
}

func SaveUserAuthInfo(user_auth_info *UserAuthInfo) {
	w := protocol.NewWriter()
	user_auth_info.Marshal(&w)
	err := db.Put([]byte(user_auth_info.AccountName), w.GetFullBytes(), nil)
	if err != nil {
		panic(err)
	}
}

func GetCurrentUIDIndex() int64 {
	b, err := db.Get([]byte("__UIDTotal"), nil)
	if err != nil {
		writer := protocol.NewWriter()
		writer.Double(0)
		db.Put([]byte("__UIDTotal"), writer.GetFullBytes(), nil)
	}
	r := protocol.NewReader(b)
	var x float64
	r.Double(&x)
	return int64(x)
}

func SetCurrentUIDIndex(ind int64) {
	w := protocol.NewWriter()
	w.Double(float64(ind))
	err := db.Put([]byte("__UIDTotal"), w.GetFullBytes(), nil)
	if err != nil {
		panic(err)
	}
}
