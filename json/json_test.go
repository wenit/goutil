package json

import (
	"testing"
	"fmt"
	"github.com/wenit/goutil/log"
)

type Person struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

type Student struct {
	Person
	Classes []string
}

func TestJsonToString(t *testing.T) {
	student := Student{Person:Person{"", 27}, Classes:[]string{"Math", "English", "Chinese"}, }
	json := JsonToString(student)
	fmt.Println(json)
}

func TestToJson(t *testing.T) {
	student := Student{}
	fmt.Println("初始：", student)
	json := `{"name":"","age":27,"Classes":["Math","English","Chinese"]}`
	ToJson(json, &student)

	fmt.Println("转换后：", student)

	logger := log.New()
	logger.Info(student)
}
