package stringutil

import (
	"github.com/satori/go.uuid"
	"time"
	"math/rand"
	"bytes"
)
//数字序列
var NUMBER_SEQ = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var CHAR_SEQ = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var NUM_CHAR_SEQ = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var ALL_CHAR_SEQ = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z"}

//@Title 获取UUID
//@Description 获取UUID
//
func UUID() string {
	return uuid.NewV4().String()
}

/**
生成指定长度的随机字符串，字符串为数字+大小写字母
 */
func RandString(length int) string {
	b := bytes.Buffer{}
	charLen := len(ALL_CHAR_SEQ)
	for i := 0; i < length; i++ {
		t := NextInt(charLen)
		s := ALL_CHAR_SEQ[t]
		b.WriteString(s)
	}
	return b.String()
}

/**
生成6位短信验证码
 */
func SmsCode() string {
	return RandNumber6()
}

/**
生成4图形验证码，数字+大写字母
 */
func ImageCode() string {
	b := bytes.Buffer{}
	charLen := len(NUM_CHAR_SEQ)
	for i := 0; i < 4; i++ {
		t := NextInt(charLen)
		s := NUM_CHAR_SEQ[t]
		b.WriteString(s)
	}
	return b.String()
}


/**
生成指定长度的随机字符串
 */
func RandNumber(length int) string {
	b := bytes.Buffer{}
	charLen := len(NUMBER_SEQ)
	for i := 0; i < length; i++ {
		t := NextInt(charLen)
		s := NUMBER_SEQ[t]
		b.WriteString(s)
	}
	return b.String()
}

func RandString6() string {
	return RandString(6)
}

func RandString4() string {
	return RandString(4)
}

func RandNumber4() string {
	return RandNumber(4)
}

func RandNumber6() string {
	return RandNumber(6)
}

func NextInt(len int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Nanosecond)
	//rand.Seed(time.Now().UnixNano())
	return r.Intn(len)
}



