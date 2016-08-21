package stringutil

import (
	"testing"
	"fmt"
	"time"
)

func TestRandString(t *testing.T) {
	l:=10
	fmt.Println(RandString(l))
	fmt.Println(RandString4())
	fmt.Println(RandString6())
}

func TestRandNumber(t *testing.T) {
	//l:=10
	fmt.Println(RandNumber4())
	fmt.Println(RandNumber6())
	startTime:=time.Now().Second()
	for i:=0;i<200 ;i++  {
		RandNumber4()
	}
	endTime:=time.Now().Second()

	fmt.Println(endTime-startTime)

}

func TestSmsCode(t *testing.T) {
	fmt.Println(SmsCode())
}

func TestImageCode(t *testing.T) {
	fmt.Println(ImageCode())
}
