package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// LogLevel 日志级别
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// 日志级别字符串
var levelStrings = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

// Logger 日志记录器
type Logger struct {
	level  LogLevel
	logger *log.Logger
	file   *os.File
}

// 默认日志记录器
var DefaultLogger *Logger

// InitLogger 初始化日志记录器
func InitLogger(level LogLevel, logFile string) error {
	// 创建日志目录
	if logFile != "" {
		logDir := filepath.Dir(logFile)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return err
		}
	}

	// 打开日志文件
	var writer io.Writer
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		// 同时输出到控制台和文件
		writer = io.MultiWriter(os.Stdout, file)
		DefaultLogger = &Logger{
			level:  level,
			logger: log.New(writer, "", 0),
			file:   file,
		}
	} else {
		// 只输出到控制台
		writer = os.Stdout
		DefaultLogger = &Logger{
			level:  level,
			logger: log.New(writer, "", 0),
		}
	}

	return nil
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// 生成日志格式
func (l *Logger) format(level LogLevel, v ...interface{}) string {
	// 获取调用者信息
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	} else {
		file = filepath.Base(file)
	}

	// 生成日志时间
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")

	// 生成日志级别
	levelStr := levelStrings[level]

	// 生成日志内容
	message := fmt.Sprint(v...)

	// 生成完整日志
	return fmt.Sprintf("%s [%s] %s:%d %s\n", timeStr, levelStr, file, line, message)
}

// Debug 调试日志
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.logger.Print(l.format(DEBUG, v...))
	}
}

// Info 信息日志
func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.logger.Print(l.format(INFO, v...))
	}
}

// Warn 警告日志
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.logger.Print(l.format(WARN, v...))
	}
}

// Error 错误日志
func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.logger.Print(l.format(ERROR, v...))
	}
}

// Fatal 致命日志
func (l *Logger) Fatal(v ...interface{}) {
	if l.level <= FATAL {
		l.logger.Print(l.format(FATAL, v...))
	}
	os.Exit(1)
}

// Debugf 调试日志格式化
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.logger.Print(l.format(DEBUG, fmt.Sprintf(format, v...)))
	}
}

// Infof 信息日志格式化
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= INFO {
		l.logger.Print(l.format(INFO, fmt.Sprintf(format, v...)))
	}
}

// Warnf 警告日志格式化
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level <= WARN {
		l.logger.Print(l.format(WARN, fmt.Sprintf(format, v...)))
	}
}

// Errorf 错误日志格式化
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.logger.Print(l.format(ERROR, fmt.Sprintf(format, v...)))
	}
}

// Fatalf 致命日志格式化
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level <= FATAL {
		l.logger.Print(l.format(FATAL, fmt.Sprintf(format, v...)))
	}
	os.Exit(1)
}

// 全局日志函数

// Debug 全局调试日志
func Debug(v ...interface{}) {
	DefaultLogger.Debug(v...)
}

// Info 全局信息日志
func Info(v ...interface{}) {
	DefaultLogger.Info(v...)
}

// Warn 全局警告日志
func Warn(v ...interface{}) {
	DefaultLogger.Warn(v...)
}

// Error 全局错误日志
func Error(v ...interface{}) {
	DefaultLogger.Error(v...)
}

// Fatal 全局致命日志
func Fatal(v ...interface{}) {
	DefaultLogger.Fatal(v...)
}

// Debugf 全局调试日志格式化
func Debugf(format string, v ...interface{}) {
	DefaultLogger.Debugf(format, v...)
}

// Infof 全局信息日志格式化
func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

// Warnf 全局警告日志格式化
func Warnf(format string, v ...interface{}) {
	DefaultLogger.Warnf(format, v...)
}

// Errorf 全局错误日志格式化
func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

// Fatalf 全局致命日志格式化
func Fatalf(format string, v ...interface{}) {
	DefaultLogger.Fatalf(format, v...)
}
