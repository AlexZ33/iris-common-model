package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Field struct {
	tableName struct{} `pg:"alias:field,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Type         string    `pg:"type,notnull" json:"type"`
	Description  string    `pg:"description" json:"description"`
	CollectionId string    `pg:"collection_id,notnull,type:uuid" json:"collection_id"`
	ReferenceId  string    `pg:"reference_id,type:uuid" json:"reference_id"`
	Equivalents  []string  `pg:"equivalents,type:uuid[]" json:"equivalents"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	Tags         []string  `pg:"tags,type:uuid[]" json:"tags"`
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

func (f *Field) Init() *Field {
	f.Id = Id()
	if f.ManagerId == "" {
		f.ManagerId = f.MaintainerId
	}
	if f.Visibility == "" {
		f.Visibility = "internal"
	}
	if f.Status == "" {
		f.Status = "active"
	}
	return f
}

func (f *Field) BasicInfo() iris.Map {
	return iris.Map{
		"id":          f.Id,
		"name":        f.Name,
		"type":        f.Type,
		"description": f.Description,
		"visibility":  f.Visibility,
		"status":      f.Status,
		"version":     f.Version,
	}
}
