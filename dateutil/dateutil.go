package dateutil

import (
	"time"
	"fmt"
)
//模板格式：2006-01-02
const DATE_FORMAT_LAYOUT = "2006-01-02"
//模板格式：20060102
const DATE_FORMAT_LAYOUT1 = "20060102"
//模板格式：2006/01/02
const DATE_FORMAT_LAYOUT2 = "2006/01/02"
//模板格式：2006/1/2
const DATE_FORMAT_LAYOUT3 = "2006/1/2"
//模板格式：2006年1月2日，中文日期
const DATE_FORMAT_LAYOUT_CN = "2006年1月2日"
//模板格式：15:04:05
const TIME_FORMAT_LAYOUT = "15:04:05"
//模板格式：150405
const TIME_FORMAT_LAYOUT1 = "150405"
//模板格式：2006-01-02 15:04:05
const DATE_TIME_FORMAT_LAYOUT = "2006-01-02 15:04:05"

/**
获取当前日期
转换模板：2006-01-02
@return	string
 */
func GetDate() string {
	return GetDateByLayout(DATE_FORMAT_LAYOUT)
}

/**
获取当前时间
转换模板：15:04:05
@return	string
 */
func GetTime() string {
	return GetDateByLayout(TIME_FORMAT_LAYOUT)
}


/**
获取当前日期+时间
转换模板：2006-01-02 15:04:05
@return	string
 */
func GetDateTime() string {
	return GetDateByLayout(DATE_TIME_FORMAT_LAYOUT)
}

/**
获取指定格式的当前日期字符串
@param	layout	日期模板
@return	string
 */
func GetDateByLayout(layout string) string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	_date := tm.Format(layout)
	return _date
}

/**
将日期字符串转换成time类型
转换模板："2006-01-02"
@param	value	待转换的时间字符串
@return	Time	时间类型
 */
func ParseDate(value string) time.Time {
	return ParseByLayout(value, DATE_FORMAT_LAYOUT)
}

/**
将时间字符串转换成time类型
转换模板："15:04:05"
@param	value	待转换的时间字符串
@return	Time	时间类型
 */
func ParseTime(value string) time.Time {
	return ParseByLayout(value, TIME_FORMAT_LAYOUT)
}

/**
将日期时间字符串转换成time类型
转换模板：2006-01-02 15:04:05
@param	value	待转换的时间字符串
@return	Time	时间类型
 */
func ParseDateTime(value string) time.Time {
	return ParseByLayout(value, DATE_TIME_FORMAT_LAYOUT)
}

/**
根据指定类型的模板转换时间字符串为时间类型
@param	value	待转换的时间字符串
@param	layout	转换模板
@return	Time	时间类型
 */
func ParseByLayout(value string, layout string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

/**
将特定日期按指定格式进行转换
@param	t	待转换的时间类型
@param	layout	转换模板
@return	Time	时间字符串
 */
func FormatDate(t time.Time, layout string) string {
	tm := time.Unix(t.Unix(), 0)
	_date := tm.Format(layout)
	return _date
}

/**
日期格式转换
@param	value	待转换的时间字符串
@param	src	源转换模板
@param	dest	目标转换模板
@return	string	目标日期字符串
 */
func FormatConvert(date string, src string, dest string) string {
	return FormatDate(ParseByLayout(date, src), dest)
}

/**
获取当前日期几天前的日期
日期模板：2006-01-02
@param	days	几天
@return	string	几天前的日期字符串
 */
func BeforeDay(days int) string {
	now := time.Now()
	return FormatDate(now.AddDate(0, 0, -days), DATE_FORMAT_LAYOUT)
}

/**
获取前日期几天后的日期
日期模板：2006-01-02
@param	days	几天
@return	string	几天后的日期字符串
 */
func AfterDay(days int) string {
	now := time.Now()
	return FormatDate(now.AddDate(0, 0, days), DATE_FORMAT_LAYOUT)
}