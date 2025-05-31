package player_store

import (
	"MineArcade-backend/minearcade-server/defines"
	"MineArcade-backend/minearcade-server/protocol"

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

func ReadPlayerStore(playerUID string) *PlayerStore {
	inf, err := OpenPlayerStoreDB().Get([]byte(playerUID), nil)
	if err != nil {
		return NewPlayerStore()
	} else {
		reader := protocol.NewReader(inf)
		playerStore := &PlayerStore{}
		playerStore.Unmarshal(&reader)
		return playerStore
	}
}

func SavePlayerStore(playerUID string, playerStore *PlayerStore) {
	writer := protocol.NewWriter()
	playerStore.Marshal(&writer)
	OpenPlayerStoreDB().Put([]byte(playerUID), writer.GetFullBytes(), nil)
}
