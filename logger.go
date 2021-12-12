package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logs struct {
	log *zap.Logger
}

func NewLogs(level string) *Logs {
	Level = level
	return &Logs{
		log: Zap(),
	}
}

func (l *Logs) writeMsg(ctx *gin.Context, msg string, payload ...interface{}) string {
	logID := GetLogID(ctx)
	errMsg := fmt.Sprintf("[%s] %+v, %+v", logID, msg, payload)
	return errMsg
}
func (l *Logs) writefMsg(ctx *gin.Context, format string, a ...interface{}) string {
	logID := GetLogID(ctx)
	return "[" + logID + "]" + fmt.Sprintf(format, a)
}

// ---- error

// Error error
func (l *Logs) Error(msg string, err error) {
	l.log.Error(msg, zap.Any("err", err))
}

// Errorf  error format
func (l *Logs) Errorf(format string, a ...interface{}) {
	l.log.Error(fmt.Sprintf(format, a))
}

// CtxError context error
func (l *Logs) CtxError(ctx *gin.Context, msg string, err error, payload ...interface{}) {
	log := Zap()
	log.Error(l.writeMsg(ctx, msg, payload), zap.Any("err", err))
}

// CtxErrorf context error format
func (l *Logs) CtxErrorf(ctx *gin.Context, format string, a ...interface{}) {
	l.log.Error(l.writefMsg(ctx, format, a))
}

// ---- info

// Info info
func (l *Logs) Info(msg string, data interface{}) {
	l.log.Info(msg, zap.Any("info", data))
}

// Infof  info format
func (l *Logs) Infof(format string, a ...interface{}) {
	l.log.Info(fmt.Sprintf(format, a))
}

// CtxInfo context info
func (l *Logs) CtxInfo(ctx *gin.Context, msg string, data interface{}, payload ...interface{}) {
	log := Zap()
	log.Info(l.writeMsg(ctx, msg, payload), zap.Any("info", data))
}

// CtxInfof context info format
func (l *Logs) CtxInfof(ctx *gin.Context, format string, a ...interface{}) {
	l.log.Info(l.writefMsg(ctx, format, a))
}

// ---- Warn

// Warn warn
func (l *Logs) Warn(msg string, data interface{}) {
	l.log.Warn(msg, zap.Any("warn", data))
}

// Warnf info format
func (l *Logs) Warnf(format string, a ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, a))
}

// CtxWarn context info
func (l *Logs) CtxWarn(ctx *gin.Context, msg string, data interface{}, payload ...interface{}) {
	l.log.Warn(l.writeMsg(ctx, msg, payload), zap.Any("warn", data))
}

// CtxWarnf context info format
func (l *Logs) CtxWarnf(ctx *gin.Context, format string, a ...interface{}) {
	l.log.Info(l.writefMsg(ctx, format, a))
}
