package event

import (
	"github.com/syafiqparadisam/mytelebot/entity"
)

func (e *event) CalculatePriceLevelDistro(os entity.Os) int64 {
	level := os.Level
	switch level {
	case "veryeasy":
		return 20000
	case "easy":
		return 30000
	case "medium":
		return 40000
	case "hard":
		return 60000
	case "veryhard":
		return 100000
	}
	return 0
}
