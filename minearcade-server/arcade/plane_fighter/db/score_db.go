package db

import (
	"MineArcade-backend/minearcade-server/defines"

	"github.com/df-mc/goleveldb/leveldb"
)

var db *leveldb.DB

func OpenScoreDB() *leveldb.DB {
	if db == nil {
		ldb, err := leveldb.OpenFile(defines.PLANE_FIGHTER_DB_PATH, nil)
		if err != nil {
			panic(err)
		}
		db = ldb
	}
	return db
}

func GetScore(playerUID string) (*PlayerScore, error) {
	dat, err := OpenScoreDB().Get([]byte(playerUID), nil)
	if err != nil {
		return nil, err
	}
	score := &PlayerScore{}
	err = score.Unmarshal(dat)
	return score, err
}
