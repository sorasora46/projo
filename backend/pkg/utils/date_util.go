package utils

import (
	"time"

	"github.com/sorasora46/projo/backend/pkg/constants"
)

func GetEpochXHoursFromNow(x uint) int64 {
	return time.Now().Add(time.Hour * time.Duration(x)).Unix()
}

func GetEpochXDaysFromNow(x uint) int64 {
	return time.Now().Add(time.Hour * constants.OneDayInHour * time.Duration(x)).Unix()
}
