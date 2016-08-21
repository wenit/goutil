package config

import (
	"testing"
	"fmt"
	"time"
	"reflect"
	"github.com/wenit/goutil/json"
	"github.com/wenit/goutil/log"
)

func TestLoad(t *testing.T) {

	path := "d:/test.ini"

	p := &Property{}
	p.Load(path)
	fmt.Println(p.Store)

	sleepTime := time.Second * 1
	fmt.Println("休眠：", sleepTime)
	time.Sleep(sleepTime)

	fmt.Println("重新加载配置文件")

	p.reload()
	fmt.Println(p.Store)

	v1 := p.Get("name", 1)

	fmt.Println(v1 == "", "type:", reflect.TypeOf(v1))

	v2 := p.GetBool("bool")

	v3 := p.GetInt("int")

	fmt.Println(v2 == false, "type:", reflect.TypeOf(v2))

	fmt.Println(v3 == 100, "type:", reflect.TypeOf(v3))

	fmt.Println(p)

	fmt.Println(json.JsonToString(p))

	logger := log.New()

	logger.Info(*p)
}
