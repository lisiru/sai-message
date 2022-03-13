package util

import (
	"strconv"
	"time"
)

func TimeStrToTime(timeStr string) time.Time  {
	loc,_:=time.LoadLocation("Local")
	theTime,_:=time.ParseInLocation("2006-01-02 15:04:05",timeStr,loc)
	return theTime


}

func SliceContains(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}


func StringToInt(str string) int {
	ret, err := strconv.Atoi(str)
	if err != nil {
		ret = 0
	}
	return ret
}

func IntToString(str int) string {
	return strconv.Itoa(str)
}

func StringToInt64(str string) int64 {
	ret, err := strconv.Atoi(str)
	if err != nil {
		ret = 0
	}
	return int64(ret)
}

func StringToUint64(str string) uint64  {
	ret, err := strconv.Atoi(str)
	if err != nil {
		ret = 0
	}
	return uint64(ret)
}

func Int64ToString(num int64) string  {
	return strconv.FormatInt(num,10)
}



