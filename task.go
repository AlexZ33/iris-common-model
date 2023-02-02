package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Task struct {
	tableName struct{} `pg:"alias:task,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Type         string    `pg:"type,notnull" json:"type"`
	Description  string    `pg:"description" json:"description"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	InputId      string    `pg:"input_id,type:uuid" json:"input_id"`
	OutputId     string    `pg:"output_id,type:uuid" json:"output_id"`
	Dependencies []string  `pg:"dependencies,type:uuid[]" json:"dependencies"`
	CustomerId   string    `pg:"customer_id,type:uuid" json:"customer_id"`
	ValidFrom    time.Time `pg:"valid_from,default:now()" json:"valid_from"`
	ExpiresAt    time.Time `pg:"expires_at,default:now()" json:"expires_at"`
	Priority     uint64    `pg:"priority,default:0" json:"priority"`
	Schedule     string    `pg:"schedule" json:"schedule"`
	LastTime     time.Time `pg:"last_time,default:'epoch'" json:"last_time"`
	NextTime     time.Time `pg:"next_time,default:'epoch'" json:"next_time"`
	ManagerId    string    `pg:"manager_id,notnull,type:uuid" json:"manager_id"`
	Visibility   string    `pg:"visibility,default:'internal'" json:"visibility"`
	Status       string    `pg:"status,default:'active'" json:"status"`
	Metrics      iris.Map  `pg:"metrics,type:jsonb" json:"metrics"`
	OtherInfo    iris.Map  `pg:"other_info,type:jsonb" json:"other_info"`
	CreatedAt    time.Time `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt    time.Time `pg:"updated_at,default:now()" json:"updated_at"`
	MaintainerId string    `pg:"maintainer_id,notnull,type:uuid" json:"maintainer_id"`
	Version      uint64    `pg:"version,default:0" json:"version"`
}

func (t *Task) Init() *Task {
	t.Id = helper.Id()
	t.Metrics = iris.Map{
		"output_snapshot":            nil,
		"running_count":              0,
		"execution_count":            0,
		"fulfill_count":              0,
		"fail_count":                 0,
		"ignore_count":               0,
		"invalid_count":              0,
		"success_count":              0,
		"running_time_ms":            0,
		"execution_time_ms":          0,
		"output_time_ms":             0,
		"completion_time_ms":         0,
		"average_running_time_ms":    0,
		"average_execution_time_ms":  0,
		"average_output_time_ms":     0,
		"average_completion_time_ms": 0,
		"max_running_time_ms":        0,
		"max_execution_time_ms":      0,
		"max_output_time_ms":         0,
		"max_completion_time_ms":     0,
	}
	if t.ManagerId == "" {
		t.ManagerId = t.MaintainerId
	}
	if t.Visibility == "" {
		t.Visibility = "internal"
	}
	if t.Status == "" {
		t.Status = "inactive"
	}
	return t
}

func (t *Task) BasicInfo() iris.Map {
	return iris.Map{
		"id":          t.Id,
		"name":        t.Name,
		"type":        t.Type,
		"description": t.Description,
		"visibility":  t.Visibility,
		"status":      t.Status,
		"version":     t.Version,
	}
}
