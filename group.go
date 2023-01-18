package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Group struct {
	Id           string    `json:"id"`
	Name         string    `json: "name"`
	Members      []string  `json:"members"`
	Description  string    `json:"description"`
	AccessKey    string    `json: "access_key"`
	ManagerId    string    `json:"manager_id"`
	Visibility   string    `json:"visibility"`
	Status       string    `json:"status"`
	Metrics      iris.Map  `json:"metrics"`
	BaseInfo     iris.Map  `json: "base_info"`
	OtherInfo    iris.Map  `json: "other_info"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json: "updated_at"`
	MaintainerId string    `json: "maintainer_id"`
	Version      uint64    `json:"version"`
}

func (g *Group) Init() *Group {
	g.Id = Id()
	g.AccessKey = AccessKey()
	if g.ManagerId == "" {
		g.ManagerId = g.MaintainerId
	}
	if g.Visibility == "" {
		g.Visibility = "internal"
	}

	if g.Status == "" {
		g.Status = "active"
	}
	return g
}

func (g *Group) BasicInfo() iris.Map {
	return iris.Map{
		"id":          g.Id,
		"name":        g.Name,
		"description": g.Description,
		"visibility":  g.Visibility,
		"status":      g.Status,
		"version":     g.Version,
	}
}

func (g *Group) GetMap() iris.Map {
	return iris.Map{
		"id":            g.Id,
		"name":          g.Name,
		"members":       g.Members,
		"description":   g.Description,
		"manager_id":    g.ManagerId,
		"visibility":    g.Visibility,
		"status":        g.Status,
		"other_info":    g.OtherInfo,
		"created_at":    g.CreatedAt,
		"updated_at":    g.UpdatedAt,
		"maintainer_id": g.MaintainerId,
		"version":       g.Version,
	}
}
