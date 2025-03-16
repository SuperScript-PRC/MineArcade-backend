package defines

import "path/filepath"

const MINEARCADE_VERSION = 100

const BASE_PATH = ".mine_arcade"
const SUB_DIR_ACCOUNT = "accounts"
const SUB_DIR_STORE = "player_store"
const SUB_DIR_MINEAREA = "public_minearea"

const SUB_FILE_ACCOUNT_DB = "accounts.db"
const SUB_FILE_STORE_DB = "player_store.db"
const SUB_FILE_MINEAREA_MAP_DATA = "map.dat"

var ACCOUNT_DIR_PATH = filepath.Join(BASE_PATH, SUB_DIR_ACCOUNT)
var PLAYER_STORE_DIR_PATH = filepath.Join(BASE_PATH, SUB_DIR_STORE)
var ACCOUNT_DB_PATH = filepath.Join(ACCOUNT_DIR_PATH, SUB_FILE_ACCOUNT_DB)
var PLAYER_STORE_DB_PATH = filepath.Join(PLAYER_STORE_DIR_PATH, SUB_FILE_STORE_DB)
var MAP_DIR_PATH = filepath.Join(BASE_PATH, SUB_DIR_MINEAREA)
var MAP_PATH = filepath.Join(MAP_DIR_PATH, SUB_FILE_MINEAREA_MAP_DATA)
