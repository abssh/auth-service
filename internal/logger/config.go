package logger

type LogConfig interface {
	GetLogLevel() LogLevel
}