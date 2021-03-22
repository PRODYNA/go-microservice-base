package log

import (
	"context"
	"fmt"
	"github.com/prodyna/go-microservice-base/trace"
	"reflect"
	"strings"
)

type LogLevel int

func (l LogLevel) string() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case TRACE:
		return "TRACE"
	default:
		return "INFO"
	}
}

const (
	TRACE LogLevel = 0
	DEBUG LogLevel = 1
	INFO  LogLevel = 2
	ERROR LogLevel = 3
)

type LogType int

const (
	PLAIN LogType = 0
	JSON  LogType = 1
)

type Logger struct {
	Level   LogLevel
	Path    string
	LogType LogType
}

type logmessage struct {
	context   context.Context
	path      string
	level     string
	start     string
	duration  string
	operation string
	message   string
	fields    map[string]string
	logtype   string
}

func RootLoggerCtx(id string, level LogLevel) *Logger {
	return &Logger{
		Level:   level,
		Path:    id,
		LogType: JSON,
	}
}

func RootLogger(id string, level LogLevel) *Logger {
	return &Logger{
		Level:   level,
		Path:    id,
		LogType: JSON,
	}
}

func (l *Logger) Logger(logger *Logger, id string) *Logger {
	return &Logger{
		Level:   logger.Level,
		Path:    logger.Path + "." + id,
		LogType: l.LogType,
	}
}

func (l *Logger) Error(operation, message string, i ...interface{}) {

	msg := logmessage{
		path:      l.Path,
		level:     ERROR.string(),
		operation: operation,
		message:   message,
		logtype:   "log",
	}

	l.log(msg, i)
}

func (l *Logger) Info(operation, message string, i ...interface{}) {
	if l.Level >= INFO {
		msg := logmessage{
			path:      l.Path,
			level:     INFO.string(),
			operation: operation,
			message:   message,
			logtype:   "log",
		}

		l.log(msg, i)
	}
}

func (l *Logger) Trace(operation, message string, i ...interface{}) {
	if l.Level >= TRACE {
		msg := logmessage{
			path:      l.Path,
			level:     TRACE.string(),
			operation: operation,
			message:   message,
			logtype:   "log",
		}

		l.log(msg, i)
	}
}

func (l *Logger) Debug(operation, message string, i ...interface{}) {
	if l.Level >= DEBUG {
		msg := logmessage{
			path:      l.Path,
			level:     DEBUG.string(),
			operation: operation,
			message:   message,
			logtype:   "log",
		}

		l.log(msg, i)
	}
}

func (l *Logger) Audit(operation, message string, i ...interface{}) {
	msg := logmessage{
		path:      l.Path,
		level:     INFO.string(),
		operation: operation,
		message:   message,
		logtype:   "audit",
	}

	l.log(msg, i)
}

func (l *Logger) log(msg logmessage, data []interface{}) {

	for _, i := range data {

		switch i.(type) {
		case context.Context:
			{
				msg.context = i.(context.Context)
				continue
			}
		}

		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Map:
			{
				msg.fields = i.(map[string]string)
			}
		}

	}

	if l.LogType == JSON {
		logJson(msg)
	} else {
		logPlain(msg)
	}

}

func fillFields(msg logmessage, i interface{}) {
	m := i.(map[string]string)
	msg.fields = m
}

func fillGenericFields(msg logmessage, i interface{}) {
	m := i.(map[string]interface{})
	msg.fields = make(map[string]string)

	for k, v := range m {
		msg.fields[k] = fmt.Sprintf("%s", v)
	}
}

func logJson(msg logmessage) {
	sb := strings.Builder{}
	sb.WriteString("{")
	sb.WriteString("\"level\": \"")
	sb.WriteString(msg.level)
	sb.WriteString("\", \"path\": \"")
	sb.WriteString(msg.path)
	sb.WriteString("\", \"op\": \"")
	sb.WriteString(msg.operation)
	sb.WriteString("\", \"msg\": \"")
	sb.WriteString(msg.message)
	sb.WriteString("\", \"type\": \"")
	sb.WriteString(msg.logtype)
	sb.WriteString("\"")
	sb.WriteString(logMsgMap("fields", msg.fields))
	if msg.context != nil && msg.context.Value(trace.TraceContextKey) != nil {
		tc := msg.context.Value(trace.TraceContextKey).(trace.TraceContext)
		sb.WriteString(", \"trace\" : { ")
		sb.WriteString("\"TraceId\": \"")
		sb.WriteString(tc.TraceId)
		sb.WriteString("\", \"SpanId\": \"")
		sb.WriteString(tc.SpanId)
		sb.WriteString("\", \"ParentSpanId\": \"")
		sb.WriteString(tc.ParentSpanId)
		sb.WriteString("\" }")
	}

	sb.WriteString("}\n")

	if msg.level == "ERROR" {
		print(sb.String())
	} else {
		fmt.Print(sb.String())
	}

}

func logMsgMap(id string, data map[string]string) string {
	sb := strings.Builder{}
	if data != nil && len(data) > 0 {
		sb.WriteString(", \"" + id + "\": { ")
		s := len(data)
		i := 1
		for k, v := range data {
			sb.WriteString("\"")
			sb.WriteString(k)
			sb.WriteString("\": \"")
			sb.WriteString(v)
			if s == i {
				sb.WriteString("\" ")
			} else {
				sb.WriteString("\", ")
			}
			i++
		}
		sb.WriteString("} ")
		return sb.String()
	} else {
		return ""
	}
}

func logPlain(msg logmessage) {
	fmt.Printf("%s %s", msg.level, msg.path)
}
