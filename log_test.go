package logger

import (
	"errors"
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLogs("debug")
	fmt.Println(log)
	log.Info("data: %+v", errors.New("hhhh"))
}
