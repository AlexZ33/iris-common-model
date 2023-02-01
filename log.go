package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"net"
	"strings"
	"time"
)

type Log struct {
	Id           string    `json:"id"`
	Service      string    `json:"service"`
	ServerHost   string    `json:"server_host"`
	ClientIP     net.IP    `json:"client_ip"`
	RecordedAt   time.Time `json:"recorded_at"`
	Level        string    `json:"level"`
	Topic        string    `json:"topic"`
	Message      string    `json:"message"`
	Content      iris.Map  `json:"content"`
	Source       string    `json:"source"`
	Visibility   string    `json:"visibility"`
	Status       string    `json:"status"`
	Metrics      iris.Map  `json:"metrics"`
	OtherInfo    iris.Map  `json:"other_info"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MaintainerId string    `json:"maintainer_id"`
	Version      uint64    `json:"version"`
}

func (l *Log) Init() *Log {
	l.Id = helper.Id()
	if l.Visibility == "" {
		l.Visibility = "internal"
	}
	if l.Status == "" {
		l.Status = "active"
	}
	return l
}

func (l *Log) Normalize() *Log {
	if level := l.Level; level != "" {
		level := strings.ToLower(strings.Trim(level, "[]"))
		if level == "warning" {
			level = "warn"
		}
		l.Level = level
	}
	if content := l.Content; len(content) > 0 {
		if method, ok := content["method"]; ok {
			content["request_method"] = method
		}
		if requestMethod, ok := content["request_method"]; ok {
			requestMethod := helper.ParseString(requestMethod)
			content["request_method"] = strings.ToUpper(requestMethod)
		}
		if statusCode, ok := content["status_code"]; ok {
			content["status_code"] = helper.ParseUint64(statusCode)
		}
		if contentLength, ok := content["content_length"]; ok {
			content["content_length"] = helper.ParseUint64(contentLength)
		}
		if responseTime, ok := content["response_time"]; ok {
			content["response_time_ms"] = helper.ParseMilliseconds(responseTime)
		}
		l.Content = content
	}
	if source := l.Source; source != "" {
		l.Source = strings.TrimPrefix(source, "file:///")
	}
	return l
}

func (l *Log) BasicInfo() iris.Map {
	return iris.Map{
		"id":         l.Id,
		"service":    l.Service,
		"level":      l.Level,
		"visibility": l.Visibility,
		"status":     l.Status,
		"version":    l.Version,
	}
}
