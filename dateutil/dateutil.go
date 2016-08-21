package dateutil

import "time"

const DATE_FORMAT_LAYOUT  = "2006-01-02"
const TIME_FORMAT_LAYOUT  = "15:04:05"
const DATE_TIME_FORMAT_LAYOUT  = "2006-01-02 15:04:05"

func GetDate() string {
	timestamp:=time.Now().Unix()
	tm:=time.Unix(timestamp,0)
	_date:=tm.Format(DATE_FORMAT_LAYOUT)
	return _date
}

func GetTime() string {
	timestamp:=time.Now().Unix()
	tm:=time.Unix(timestamp,0)
	_date:=tm.Format(TIME_FORMAT_LAYOUT)
	return _date
}

func GetDateTime() string {
	timestamp:=time.Now().Unix()
	tm:=time.Unix(timestamp,0)
	_date:=tm.Format(DATE_TIME_FORMAT_LAYOUT)
	return _date
}