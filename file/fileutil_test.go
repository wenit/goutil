package file

import (
	"testing"
	"fmt"
)

func TestAppAbsPath(t *testing.T) {
	fmt.Println(AppAbsPath())
}


func TestAppDir(t *testing.T) {
	fmt.Println(AppDir())
}
