package maputil

import (
	"testing"
	"fmt"
)

func TestGetValue(t *testing.T) {
	m:=make(map[string]interface{})
	m["a"]=1
	fmt.Println(GetValue(m,"b",2))
}

func TestGetIntValue(t *testing.T) {
	m:=make(map[string]interface{})
	m["b"]=10
	fmt.Println(GetIntValue(m,"a"))

	m["b"]="20"
	fmt.Println(GetIntValue(m,"b"))
	fmt.Println(GetIntValue(m,"a",10))
}

func TestGetFloatValue(t *testing.T) {
	m:=make(map[string]interface{})
	m["a"]=1000
	fmt.Println(GetFloatValue(m,"b"))
}

func TestGetBoolValue(t *testing.T) {
	m:=make(map[string]interface{})
	m["a"]=true
	fmt.Println(GetBoolValue(m,"a",true))
}

func TestName(t *testing.T) {
	var v interface{}
	v=112
	fmt.Println(string(v.(int)))
}

func TestGetStringValue(t *testing.T) {
	m:=make(map[string]interface{})
	m["b"]="abc"
	fmt.Println(GetStringValue(m,"a","bb"))
}
