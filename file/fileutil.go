package file

import (
	"os/exec"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
)

/*
获取应用绝对路径
@return string
 */
func AppAbsPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

/**
获取应用目录
@return string
 */
func AppDir() string {
	dir := filepath.Dir(AppAbsPath())
	return dir
}

/**
读取文件，返回文件内容字符串
@param	fileName	文件名称
@return	string
 */
func ReadFile(fileName string) string {
	bytes, _ := ioutil.ReadFile(fileName)
	return string(bytes)
}

/**
获取文件后缀名，不包含小数点
@param	fileName	文件名称
@return	string
 */
func Ext(fileName string) string {
	return strings.Replace(filepath.Ext(fileName), ".", "", 1)
}

/**
像文件中写入数据，如果文件不存在，会创建文件，文件存在会覆盖文件中的数据
@param	fileName	文件路径
@param	data		数据
 */
func WriteFile(fileName string, data string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0x644)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}

	return err
}

/**
追加写文件，向文件的末尾追加写入数据
@param	fileName	文件路径
@param	data		数据
 */
func AppendFile(fileName string, data string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_RDWR, os.ModePerm | os.ModeTemporary)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return err
}