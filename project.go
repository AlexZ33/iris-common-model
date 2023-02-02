package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Project struct {
	tableName struct{} `pg:"alias:project,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Description  string    `pg:"description" json:"description"`
	ParentId     string    `pg:"parent_id,type:uuid" json:"parent_id"`
	Tenants      []string  `pg:"tenants,type:uuid[]" json:"tenants"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
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

func (p *Project) Init() *Project {
	p.Id = helper.Id()
	if p.ManagerId == "" {
		p.ManagerId = p.MaintainerId
	}
	if p.Visibility == "" {
		p.Visibility = "internal"
	}
	if p.Status == "" {
		p.Status = "active"
	}
	return p
}

func (p *Project) BasicInfo() iris.Map {
	return iris.Map{
		"id":          p.Id,
		"name":        p.Name,
		"description": p.Description,
		"visibility":  p.Visibility,
		"status":      p.Status,
		"version":     p.Version,
	}
}
