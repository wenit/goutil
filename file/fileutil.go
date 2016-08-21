package file

import (
	"os/exec"
	"os"
	"path/filepath"
)

/*
获取应用路径
 */
func AppAbsPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path

}


/**
获取应用目录
 */
func AppDir() string {
	dir := filepath.Dir(AppAbsPath())
	return dir
}
