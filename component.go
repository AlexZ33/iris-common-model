package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Component struct {
	tableName struct{} `pg:"alias:component,discard_unknown_columns"`

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

func (c *Component) Init() *Component {
	c.Id = helper.Id()
	if c.ManagerId == "" {
		c.ManagerId = c.MaintainerId
	}
	if c.Visibility == "" {
		c.Visibility = "internal"
	}
	if c.Status == "" {
		c.Status = "active"
	}
	return c
}

func (c *Component) BasicInfo() iris.Map {
	return iris.Map{
		"id":          c.Id,
		"name":        c.Name,
		"description": c.Description,
		"visibility":  c.Visibility,
		"status":      c.Status,
		"version":     c.Version,
	}
}
