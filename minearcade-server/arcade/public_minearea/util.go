package public_minearea

import "time"

func float64time() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}
