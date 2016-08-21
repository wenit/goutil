package dateutil

import (
	"testing"
	"fmt"
)

func TestGetDate(t *testing.T) {
	fmt.Println(GetDate())
}

func TestGetTime(t *testing.T) {
	fmt.Println(GetTime())
}

func TestGetDateTime(t *testing.T) {
	fmt.Println(GetDateTime())
}
