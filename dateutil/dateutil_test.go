package dateutil

import (
	"testing"
	"fmt"
)

func TestGetDateByLayout(t *testing.T) {
	fmt.Println(GetDateByLayout("2006/01/02 15:04:05"))
	//output:2016/08/22 13:13:25
}

func TestGetDate(t *testing.T) {
	fmt.Println(GetDate())
	//output:2016-08-22
}

func TestGetTime(t *testing.T) {
	fmt.Println(GetTime())
	//output:13:13:25
}

func TestGetDateTime(t *testing.T) {
	fmt.Println(GetDateTime())
	//output:2016-08-22 13:13:25
}

func TestParseTime(t *testing.T) {
	var time1="11:45:21"
	tm:=ParseTime(time1)
	fmt.Println(tm)
	//output:0000-01-01 11:45:21 +0800 CST

	time1="2016-08-21"
	tm=ParseDate(time1)
	fmt.Println(tm)
	//output:2016-08-21 00:00:00 +0800 CST

	time1="2016-08-21 11:45:21"
	tm=ParseDateTime(time1)
	fmt.Println(tm)
	//output:2016-08-21 11:45:21 +0800 CST

	time1="2016/8/21 1:45:23"
	tm=ParseByLayout(time1,"2006/1/02 15:04:05")
	fmt.Println(tm)
	//output:2016-08-21 01:45:23 +0800 CST
}