package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Form struct {
	tableName struct{} `pg:"alias:form,discard_unknown_columns"`

	Id           string     `pg:"id,pk,type:uuid" json:"id"`
	Name         string     `pg:"name,notnull" json:"name"`
	Description  string     `pg:"description" json:"description"`
	Rules        []iris.Map `pg:"rules,type:jsonb" json:"rules"`
	DatasetId    string     `pg:"dataset_id,notnull,type:uuid" json:"dataset_id"`
	CustomerId   string     `pg:"customer_id,type:uuid" json:"customer_id"`
	ValidFrom    time.Time  `pg:"valid_from,default:now()" json:"valid_from"`
	ExpiresAt    time.Time  `pg:"expires_at,default:now()" json:"expires_at"`
	Content      iris.Map   `pg:"content,type:jsonb" json:"content"`
	ManagerId    string     `pg:"manager_id,notnull,type:uuid" json:"manager_id"`
	Visibility   string     `pg:"visibility,default:'internal'" json:"visibility"`
	Status       string     `pg:"status,default:'active'" json:"status"`
	Metrics      iris.Map   `pg:"metrics,type:jsonb" json:"metrics"`
	OtherInfo    iris.Map   `pg:"other_info,type:jsonb" json:"other_info"`
	CreatedAt    time.Time  `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt    time.Time  `pg:"updated_at,default:now()" json:"updated_at"`
	MaintainerId string     `pg:"maintainer_id,notnull,type:uuid" json:"maintainer_id"`
	Version      uint64     `pg:"version,default:0" json:"version"`
}

func (f *Form) Init() *Form {
	f.Id = helper.Id()
	if f.ManagerId == "" {
		f.ManagerId = f.MaintainerId
	}
	if f.Visibility == "" {
		f.Visibility = "internal"
	}
	if f.Status == "" {
		f.Status = "active"
	}
	return f
}

func (f *Form) BasicInfo() iris.Map {
	return iris.Map{
		"id":          f.Id,
		"name":        f.Name,
		"description": f.Description,
		"visibility":  f.Visibility,
		"status":      f.Status,
		"version":     f.Version,
	}
}
