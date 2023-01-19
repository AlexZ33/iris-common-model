package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Collection struct {
	Id           string    `json:"id"`
	Name         string    `json: "name"`
	Type         string    `json:"type"`
	Description  string    `json:"description"`
	RepositoryId string    `json:"repository_id"`
	Content      iris.Map  `json:"content"`
	Tags         []string  `json:"tags"`
	CustomerId   string    `json:"customer_id"`
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
