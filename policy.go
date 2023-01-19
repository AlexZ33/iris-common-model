package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Policy struct {
	tableName struct{} `pg:"alias:policy,discard_unknown_columns"`

	Id           string    `json:"id"`
	Name         string    `json:"name"`
	TenantId     string    `json:"tenant_id"`
	Subjects     []string  `json:"subjects"`
	Resources    []string  `json:"resources"`
	Description  string    `json:"description"`
	Actions      []string  `json:"actions"`
	Effect       string    `json:"effect"`
	Conditions   iris.Map  `json:"conditions"`
	ValidFrom    time.Time `json:"valid_from"`
	ExpiresAt    time.Time `json:"expires_at"`
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

func (p *Policy) Init() *Policy {
	p.Id = Id()
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

func (p *Policy) BasicInfo() iris.Map {
	return iris.Map{
		"id":          p.Id,
		"name":        p.Name,
		"description": p.Description,
		"visibility":  p.Visibility,
		"status":      p.Status,
		"version":     p.Version,
	}
}

func (p *Policy) GetMap() iris.Map {
	return iris.Map{
		"id":            p.Id,
		"name":          p.Name,
		"tenant_id":     p.TenantId,
		"subjects":      p.Subjects,
		"resources":     p.Resources,
		"description":   p.Description,
		"actions":       p.Actions,
		"effect":        p.Effect,
		"conditions":    p.Conditions,
		"valid_from":    p.ValidFrom,
		"expires_at":    p.ExpiresAt,
		"manager_id":    p.ManagerId,
		"visibility":    p.Visibility,
		"status":        p.Status,
		"other_info":    p.OtherInfo,
		"created_at":    p.CreatedAt,
		"updated_at":    p.UpdatedAt,
		"maintainer_id": p.MaintainerId,
		"version":       p.Version,
	}
}
