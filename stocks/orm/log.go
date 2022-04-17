package orm

import (
	"context"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
	"ww/stocks/xlog"
)

var (
	infoStr     = "%s\n[info] "
	warnStr     = "%s\n[warn] "
	errStr      = "%s\n[error] "
	traceStr    = "[%.3fms] [rows:%v] %s"
	traceErrStr = "%s [%.3fms] [rows:%v] %s"

	DefaultLogger = &DBLogWriter{
		level: logger.Info,
	}
)

// DBLogWriter 实现logger接口
type DBLogWriter struct {
	level logger.LogLevel
	props map[string]string
}

// SetProperty 设置logger属性记录
func (l *DBLogWriter) SetProperty(key, val string) {
	if l.props == nil {
		l.props = make(map[string]string)
	}
	l.props[key] = val
}

// LogMode 改变日志等级
func (l *DBLogWriter) LogMode(level logger.LogLevel) logger.Interface {
	return &DBLogWriter{level: level}
}

// Info print info
func (l *DBLogWriter) Info(ctx context.Context, msg string, data ...interface{}) {
	fields := make(map[string]interface{})
	for k, v := range l.props {
		fields[k] = v
	}
	log := xlog.WithCataLog("orm").WithFields(fields)
	log.Infof(infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// Warn print warn messages
func (l *DBLogWriter) Warn(ctx context.Context, msg string, data ...interface{}) {
	fields := make(map[string]interface{})
	for k, v := range l.props {
		fields[k] = v
	}
	log := xlog.WithCataLog("orm").WithFields(fields)
	log.Infof(warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// Error print error messages
func (l *DBLogWriter) Error(ctx context.Context, msg string, data ...interface{}) {
	fields := make(map[string]interface{})
	for k, v := range l.props {
		fields[k] = v
	}
	log := xlog.WithCataLog("orm").WithFields(fields)
	log.Infof(errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

// Trace 打印具体SQL请求信息
func (l *DBLogWriter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	duration := float64(elapsed.Nanoseconds()) / 1e6
	sql, rows := fc()
	fields := map[string]interface{}{
		"execute_time": begin,
		"duration":     duration,
		"affected":     rows,
	}
	for k, v := range l.props {
		fields[k] = v
	}
	log := xlog.WithCataLog("orm").WithFields(fields)
	if err != nil {
		log.Errorf(traceErrStr, err, duration, rows, sql)
	} else {
		log.Infof(traceStr, duration, rows, sql)
	}
}
