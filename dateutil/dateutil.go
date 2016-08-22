package dateutil

import (
	"time"
	"fmt"
)

const DATE_FORMAT_LAYOUT  = "2006-01-02"
const TIME_FORMAT_LAYOUT  = "15:04:05"
const DATE_TIME_FORMAT_LAYOUT  = "2006-01-02 15:04:05"

/**
获取当前日期
转换模板：2006-01-02
return	string
 */
func GetDate() string {
	timestamp:=time.Now().Unix()
	tm:=time.Unix(timestamp,0)
	_date:=tm.Format(DATE_FORMAT_LAYOUT)
	return _date
}

/**
获取当前时间
转换模板：15:04:05
return	string
 */
func GetTime() string {
	timestamp:=time.Now().Unix()
	tm:=time.Unix(timestamp,0)
	_date:=tm.Format(TIME_FORMAT_LAYOUT)
	return _date
}

/**
获取当前日期+时间
转换模板：2006-01-02 15:04:05
return	string
 */
func GetDateTime() string {
	timestamp:=time.Now().Unix()
	tm:=time.Unix(timestamp,0)
	_date:=tm.Format(DATE_TIME_FORMAT_LAYOUT)
	return _date
}

/**
将日期字符串转换成time类型
转换模板："2006-01-02"
@param	value	待转换的时间字符串
@return	Time	时间类型
 */
func ParseDate(value string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t,err:=time.ParseInLocation(DATE_FORMAT_LAYOUT,value,loc)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

/**
将时间字符串转换成time类型
转换模板："15:04:05"
@param	value	待转换的时间字符串
@return	Time	时间类型
 */
func ParseTime(value string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t,err:=time.ParseInLocation(TIME_FORMAT_LAYOUT,value,loc)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

/**
将日期时间字符串转换成time类型
转换模板：2006-01-02 15:04:05
@param	value	待转换的时间字符串
@return	Time	时间类型
 */
func ParseDateTime(value string) time.Time {
	return ParseByLayout(value,DATE_TIME_FORMAT_LAYOUT)
}

/**
根据指定类型的模板转换时间字符串为时间类型
@param	value	待转换的时间字符串
@param	layout	转换模板
@return	Time	时间类型
 */
func ParseByLayout(value string ,layout string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t,err:=time.ParseInLocation(layout,value,loc)
	if err != nil {
		fmt.Println(err)
	}
	return t
}