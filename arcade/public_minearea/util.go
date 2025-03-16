package public_minearea

import "time"

func float32time() float32 {
	return float32(time.Now().UnixNano()) / 1e9
}
