package util

import "time"

func TimeStrToTime(timeStr string) time.Time  {
	loc,_:=time.LoadLocation("Local")
	theTime,_:=time.ParseInLocation("2006-01-02 15:04:05",timeStr,loc)
	return theTime


}

