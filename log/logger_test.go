package log

import (
	"testing"
	"fmt"
)

func TestInfo(t *testing.T) {
	logger := New()

	count := 1000
	for i := 0; i < count; i++ {
		logger.Debug("hello %d %s world", 1,"aa")
		logger.Info("hello %s world", "INFO")
		logger.Warn("hello %s world", "WARN")
		logger.Error("hello %s world", "ERROR")

		logger.Info("aaa", "aa", "bbb")
		logger.Info("bbb")

		fmt.Println("aaa", "aa", "bbb")
	}
}