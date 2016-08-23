package stringutil

import (
	"github.com/satori/go.uuid"
	"time"
	"math/rand"
	"bytes"
)

//数字序列
var NUMBER_SEQ = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

//大小字母序列
var CHAR_SEQ = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//数字+大写字母序列
var NUM_CHAR_SEQ = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//数字+大小写字母序列
var ALL_CHAR_SEQ = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z"}

/**
@Title 获取UUID
@Description 获取UUID
@return	string	uuid
*/
func UUID() string {
	return uuid.NewV4().String()
}

/**
生成指定长度的随机字符串，字符串为数字+大小写字母
@param	length	长度
@return	string	随机字符串
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
@return	string	6位长度的数字
 */
func SmsCode() string {
	return RandNumber6()
}

/**
生成4位图形验证码，数字+大写字母
@return	string	4位长度短信验证码
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
生成指定长度的随机数字字符串
@param	length	长度
@return	string	随机数字字符串
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

/**
生成6位长度的随机字符串
@return	string	6位长度随机字符串
 */
func RandString6() string {
	return RandString(6)
}

/**
生成4位长度的随机字符串
@return	string	4位长度随机字符串
 */
func RandString4() string {
	return RandString(4)
}

/**
生成4位长度的随机数字字符串
@return	string	4位随机数字字符串
 */
func RandNumber4() string {
	return RandNumber(4)
}

/**
生成6位长度的随机数字字符串
@return	string	6位随机数字字符串
 */
func RandNumber6() string {
	return RandNumber(6)
}

/**
随机生成指定范围内数字
@param	length	长度
@return	int	指定范围内的数字
 */
func NextInt(length int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Nanosecond)
	//rand.Seed(time.Now().UnixNano())
	return r.Intn(length)
}




