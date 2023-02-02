package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Tag struct {
	tableName struct{} `pg:"alias:tag,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Category     string    `pg:"category,notnull" json:"category"`
	Description  string    `pg:"description" json:"description"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	ParentId     string    `pg:"parent_id,type:uuid" json:"parent_id"`
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

func (t *Tag) Init() *Tag {
	t.Id = Id()
	if t.ManagerId == "" {
		t.ManagerId = t.MaintainerId
	}
	if t.Visibility == "" {
		t.Visibility = "internal"
	}
	if t.Status == "" {
		t.Status = "active"
	}
	return t
}

func (t *Tag) BasicInfo() iris.Map {
	return iris.Map{
		"id":          t.Id,
		"name":        t.Name,
		"category":    t.Category,
		"description": t.Description,
		"visibility":  t.Visibility,
		"status":      t.Status,
		"version":     t.Version,
	}
}
