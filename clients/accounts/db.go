package accounts

import (
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol"

	"github.com/df-mc/goleveldb/leveldb"
	"github.com/pterm/pterm"
)

var db *leveldb.DB

func OpenAccountDB() *leveldb.DB {
	if db == nil {
		pterm.Info.Println("正在读取账号数据库")
		ldb, err := leveldb.OpenFile(defines.ACCOUNT_DB_PATH, nil)
		if err != nil {
			panic(err)
		}
		db = ldb
	}
	if _, err := db.Get([]byte("admin"), nil); err != nil {
		ud := &UserAuthInfo{AccountName: "admin", Nickname: "admin", PasswordMD5: "1qRG9tE+TwwcupTrbKLT9AAA"}
		w := protocol.NewWriter()
		ud.Marshal(&w)
		db.Put([]byte(ud.AccountName), w.GetFullBytes(), nil)
		pterm.Info.Println("管理员账号未初始化, 已进行初始化")
	}
	if _, err := db.Get([]byte("UUIDTotal"), nil); err != nil {
		b := protocol.NewWriter()
		b.Double(0)
		db.Put([]byte("UUIDTotal"), b.GetFullBytes(), nil)
	}
	return db
}
