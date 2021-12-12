package logger

import (
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"

	"github.com/gin-gonic/gin"
)

// PathExists 判断是否有该路径
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetWriteSyncer zap logger中加入file-rotate logs
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(LogConfig.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(LogConfig.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if LogConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// GetLogID 返回context 中的 log ID
func GetLogID(ctx *gin.Context) string {
	data, ok := ctx.Get("CONTEXT_LOG")
	if !ok {
		return ""
	}
	if logID ,ok := data.(string); !ok {
		return ""
	}else {
		return logID
	}
}
