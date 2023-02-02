package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"net"
	"strings"
	"time"
)

type Log struct {
	tableName struct{} `pg:"alias:log,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Service      string    `pg:"service,notnull" json:"service"`
	ServerHost   string    `pg:"server_host,notnull" json:"server_host"`
	ClientIP     net.IP    `pg:"client_ip,notnull" json:"client_ip"`
	RecordedAt   time.Time `pg:"recorded_at,notnull" json:"recorded_at"`
	Level        string    `pg:"level" json:"level"`
	Topic        string    `pg:"topic" json:"topic"`
	Message      string    `pg:"message,notnull" json:"message"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	Source       string    `pg:"source" json:"source"`
	Visibility   string    `pg:"visibility,default:'internal'" json:"visibility"`
	Status       string    `pg:"status,default:'active'" json:"status"`
	Metrics      iris.Map  `pg:"metrics,type:jsonb" json:"metrics"`
	OtherInfo    iris.Map  `pg:"other_info,type:jsonb" json:"other_info"`
	CreatedAt    time.Time `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt    time.Time `pg:"updated_at,default:now()" json:"updated_at"`
	MaintainerId string    `pg:"maintainer_id,notnull,type:uuid" json:"maintainer_id"`
	Version      uint64    `pg:"version,default:0" json:"version"`
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
