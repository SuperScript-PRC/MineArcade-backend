package define

import "time"

const DEBUG_SHOW_BULLET = false

const MSPT = 100
const TPS = 1000 / MSPT
const SLEEP_TIME = time.Millisecond * MSPT
const STAGE_WIDTH = 1280
const STAGE_HEIGHT = 720
const STAGE_MAX_ENTITIES = 256

// 各阶段
const (
	PHASE_PREPARE = iota
	PHASE_1
	PHASE_2
	PHASE_END
)

// 每阶段所需的总时间。
const STAGE_PREPARE_REMAIN_TIME_TICKS = 6 * TPS
const STAGE_1_REMAIN_TIME_TICKS = 60 * TPS
const STAGE_2_REMAIN_TIME_TICKS = 60 * TPS
const STAGE_END_REMAIN_TIME_TICKS = 600 * TPS

// 每阶段实体生成的冷却时间。
const STAGE_1_SPAWN_TICKS = 12
const STAGE_2_SPAWN_TICKS = 5
