package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Repository struct {
	tableName struct{} `pg:"alias:repository,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Type         string    `pg:"type,notnull" json:"type"`
	Description  string    `pg:"description" json:"description"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	Tags         []string  `pg:"tags,type:uuid[]" json:"tags"`
	ProviderId   string    `pg:"provider_id,notnull,type:uuid" json:"provider_id"`
	Tenants      []string  `pg:"tenants,type:uuid[]" json:"tenants"`
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

func (r *Repository) Init() *Repository {
	r.Id = helper.Id()
	if r.ManagerId == "" {
		r.ManagerId = r.MaintainerId
	}
	if r.ProviderId == "" {
		r.ProviderId = r.ManagerId
	}
	if r.Visibility == "" {
		r.Visibility = "internal"
	}
	if r.Status == "" {
		r.Status = "active"
	}
	return r
}

func (r *Repository) BasicInfo() iris.Map {
	return iris.Map{
		"id":          r.Id,
		"name":        r.Name,
		"type":        r.Type,
		"description": r.Description,
		"visibility":  r.Visibility,
		"status":      r.Status,
		"version":     r.Version,
	}
}
