package dateutil

import (
	"testing"
	"fmt"
)

func TestGetDate(t *testing.T) {
	fmt.Println(GetDate())
}

func TestGetTime(t *testing.T) {
	fmt.Println(GetTime())
}

func TestGetDateTime(t *testing.T) {
	fmt.Println(GetDateTime())
}

func TestParseTime(t *testing.T) {
	var time1="11:45:21"
	tm:=ParseTime(time1)
	//output:0000-01-01 11:45:21 +0800 CST
	fmt.Println(tm)

	time1="2016-08-21"
	tm=ParseDate(time1)
	//output:2016-08-21 00:00:00 +0800 CST
	fmt.Println(tm)

	time1="2016-08-21 11:45:21"
	tm=ParseDateTime(time1)
	//output:2016-08-21 11:45:21 +0800 CST
	fmt.Println(tm)

	time1="2016/8/21 1:45:23"
	tm=ParseByLayout(time1,"2006/1/02 15:04:05")
	//output:2016-08-21 01:45:23 +0800 CST
	fmt.Println(tm)
}