package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

type Policy struct {
	tableName struct{} `pg:"alias:policy,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	TenantId     string    `pg:"tenant_id,notnull,type:uuid" json:"tenant_id"`
	Subjects     []string  `pg:"subjects,array" json:"subjects"`
	Resources    []string  `pg:"resources,array" json:"resources"`
	Description  string    `pg:"description" json:"description"`
	Actions      []string  `pg:"actions,array" json:"actions"`
	Effect       string    `pg:"effect,default:'allow'" json:"effect"`
	Conditions   iris.Map  `pg:"conditions,type:jsonb" json:"conditions"`
	ValidFrom    time.Time `pg:"valid_from,default:now()" json:"valid_from"`
	ExpiresAt    time.Time `pg:"expires_at,default:now()" json:"expires_at"`
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
