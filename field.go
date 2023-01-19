package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Field struct {
	tableName struct{} `pg:"alias:field,discard_unknown_columns"`

	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Description  string    `json:"description"`
	CollectionId string    `json:"collection_id"`
	ReferenceId  string    `json:"reference_id"`
	Equivalents  []string  `json:"equivalents"`
	Content      iris.Map  `json:"content"`
	Tags         []string  `json:"tags"`
	ManagerId    string    `json:"manager_id"`
	Visibility   string    `json:"visibility"`
	Status       string    `json:"status"`
	Metrics      iris.Map  `json:"metrics"`
	OtherInfo    iris.Map  `json:"other_info"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MaintainerId string    `json:"maintainer_id"`
	Version      uint64    `json:"version"`
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
