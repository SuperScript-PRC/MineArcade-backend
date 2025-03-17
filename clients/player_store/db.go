package player_store

import (
	"MineArcade-backend/defines"
	"MineArcade-backend/protocol"

	"github.com/df-mc/goleveldb/leveldb"
)

var db *leveldb.DB

func OpenPlayerStoreDB() *leveldb.DB {
	if db == nil {
		ldb, err := leveldb.OpenFile(defines.PLAYER_STORE_DB_PATH, nil)
		if err != nil {
			panic(err)
		}
		db = ldb
	}
	return db
}

func ReadPlayerStore(playerUUID string) *PlayerStore {
	inf, err := OpenPlayerStoreDB().Get([]byte(playerUUID), nil)
	if err != nil {
		return NewPlayerStore()
	} else {
		reader := protocol.NewReader(inf)
		playerStore := &PlayerStore{}
		playerStore.Unmarshal(&reader)
		return playerStore
	}
}

func SavePlayerStore(playerUUID string, playerStore *PlayerStore) {
	writer := protocol.NewWriter()
	playerStore.Marshal(&writer)
	OpenPlayerStoreDB().Put([]byte(playerUUID), writer.GetFullBytes(), nil)
}
