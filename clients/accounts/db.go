package accounts

import (
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol"

	"github.com/df-mc/goleveldb/leveldb"
)

var db *leveldb.DB

func OpenAccountDB() *leveldb.DB {
	ldb, err := leveldb.OpenFile(defines.ACCOUNT_DB_PATH, nil)
	if err != nil {
		panic(err)
	}
	db = ldb
	if _, err := db.Get([]byte("UUIDTotal"), nil); err != nil {
		b := protocol.NewWriter()
		b.Double(0)
		db.Put([]byte("UUIDTotal"), b.GetFullBytes(), nil)
	}
	return db
}
