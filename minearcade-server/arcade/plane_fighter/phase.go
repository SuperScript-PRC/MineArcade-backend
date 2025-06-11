package plane_fighter

import "MineArcade-backend/minearcade-server/arcade/plane_fighter/define"

func (s *PlaneFighterStage) addTick() {
	s.Ticks++
	s.TicksLeft--
	if s.TicksLeft <= 0 {
		switch s.Phase {
		case define.PHASE_PREPARE:
			s.Phase = define.PHASE_1
			s.TicksLeft = define.STAGE_1_REMAIN_TIME_TICKS
		case define.PHASE_1:
			s.Phase = define.PHASE_2
			s.TicksLeft = define.STAGE_2_REMAIN_TIME_TICKS
		case define.PHASE_2:
			s.Phase = define.PHASE_END
			s.TicksLeft = define.STAGE_END_REMAIN_TIME_TICKS
		}
	}
}

func (s *PlaneFighterStage) isEnded() bool {
	return s.Phase == define.PHASE_END
}
