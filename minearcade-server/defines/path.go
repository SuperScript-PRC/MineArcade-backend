package defines

import "path/filepath"

const BASE_PATH = ".mine_arcade"
const SUB_DIR_ACCOUNT = "accounts"
const SUB_DIR_STORE = "player_store"
const SUB_DIR_ARCADE = "arcade_datas"

const SUB_FILE_ACCOUNT_DB = "accounts_db"
const SUB_FILE_STORE_DB = "player_store_db"
const SUB_FILE_MINEAREA_MAP_DATA = "minearea_map.dat"
const SUB_FILE_PLANE_FIGHTER_DB = "plane_fighter_db"

var ACCOUNT_DIR_PATH = filepath.Join(BASE_PATH, SUB_DIR_ACCOUNT)
var PLAYER_STORE_DIR_PATH = filepath.Join(BASE_PATH, SUB_DIR_STORE)
var ACCOUNT_DB_PATH = filepath.Join(ACCOUNT_DIR_PATH, SUB_FILE_ACCOUNT_DB)
var PLAYER_STORE_DB_PATH = filepath.Join(PLAYER_STORE_DIR_PATH, SUB_FILE_STORE_DB)
var ARCADE_DATA_PATH = filepath.Join(BASE_PATH, SUB_DIR_ARCADE)
var MINEAREA_MAP_PATH = filepath.Join(ARCADE_DATA_PATH, SUB_FILE_MINEAREA_MAP_DATA)
var PLANE_FIGHTER_DB_PATH = filepath.Join(ARCADE_DATA_PATH, SUB_FILE_PLANE_FIGHTER_DB)
