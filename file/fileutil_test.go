package file

import (
	"testing"
	"fmt"
	"time"
)

func TestAppAbsPath(t *testing.T) {
	fmt.Println(AppAbsPath())
}


func TestAppDir(t *testing.T) {
	fmt.Println(AppDir())
}

func TestReadFile(t *testing.T) {
	fileName:="d:/test.ini"
	fmt.Println(ReadFile(fileName))
}

func TestExt(t *testing.T) {
	fmt.Println(Ext("aa.txt.xml"))
}

func TestWriteFile(t *testing.T) {
	world:="你好，世界！"
	fileName:="d:/temp/world.txt"
	WriteFile(fileName,world)
	fmt.Println(ReadFile(fileName))
}

func TestAppendFile(t *testing.T) {
	world:="Hello World！"
	fileName:="d:/temp/world.txt"
	AppendFile(fileName,world)
	fmt.Println(ReadFile(fileName))
}