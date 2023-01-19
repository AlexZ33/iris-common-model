package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Tag struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Content      iris.Map  `json:"content"`
	ParentId     string    `json:"parent_id"`
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
		"description": t.Description,
		"visibility":  t.Visibility,
		"status":      t.Status,
		"version":     t.Version,
	}
}
