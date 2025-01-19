package Game

import (
	"fmt"
	"math"
)

type Stage struct {
	RoundNumber    int
	PointsBaseGoal float64
}

func FirstStage() *Stage {
	return &Stage{1, 300}
}

func (s *Stage) String() string {
	return fmt.Sprintf("Stage: %d, Goal: %.0f", s.RoundNumber, s.PointsBaseGoal)
}

func (s *Stage) NewStage() *Stage {
	return &Stage{s.RoundNumber + 1, float64(int(math.Round(s.PointsBaseGoal * 1.5)))}
}
