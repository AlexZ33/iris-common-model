package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Collection struct {
	tableName struct{} `pg:"alias:collection,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Type         string    `pg:"type,notnull" json:"type"`
	Description  string    `pg:"description" json:"description"`
	RepositoryId string    `pg:"repository_id,notnull,type:uuid" json:"repository_id"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	Tags         []string  `pg:"tags,type:uuid[]" json:"tags"`
	CustomerId   string    `pg:"customer_id,type:uuid" json:"customer_id"`
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

func (c *Collection) Init() *Collection {
	c.Id = Id()
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

func (c *Collection) BasicInfo() iris.Map {
	return iris.Map{
		"id":          c.Id,
		"name":        c.Name,
		"type":        c.Type,
		"description": c.Description,
		"visibility":  c.Visibility,
		"status":      c.Status,
		"version":     c.Version,
	}
}
