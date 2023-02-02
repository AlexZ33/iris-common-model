package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Group struct {
	tableName struct{} `pg:"alias:group,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Members      []string  `pg:"members,type:uuid[]" json:"members"`
	Description  string    `pg:"description" json:"description"`
	AccessKey    string    `pg:"access_key,notnull" json:"access_key"`
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
