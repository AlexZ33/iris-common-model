package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type DatasetCache struct {
	Dataset   Dataset
	CreatedAt time.Time
	Count     uint64
}

type Dataset struct {
	tableName struct{} `pg:"alias:dataset,discard_unknown_columns"`

	Id            string    `pg:"id,pk,type:uuid" json:"id"`
	Name          string    `pg:"name,notnull" json:"name"`
	Type          string    `pg:"type,notnull" json:"type"`
	Description   string    `pg:"description" json:"description"`
	Content       iris.Map  `pg:"content,type:jsonb" json:"content"`
	Tags          []string  `pg:"tags,type:uuid[]" json:"tags"`
	ApplicationId string    `pg:"application_id,notnull,type:uuid" json:"application_id"`
	ValidFrom     time.Time `pg:"valid_from,default:now()" json:"valid_from"`
	ExpiresAt     time.Time `pg:"expires_at,default:now()" json:"expires_at"`
	Message       string    `pg:"message" json:"message"`
	ManagerId     string    `pg:"manager_id,notnull,type:uuid" json:"manager_id"`
	Visibility    string    `pg:"visibility,default:'internal'" json:"visibility"`
	Status        string    `pg:"status,default:'active'" json:"status"`
	Metrics       iris.Map  `pg:"metrics,type:jsonb" json:"metrics"`
	OtherInfo     iris.Map  `pg:"other_info,type:jsonb" json:"other_info"`
	CreatedAt     time.Time `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt     time.Time `pg:"updated_at,default:now()" json:"updated_at"`
	MaintainerId  string    `pg:"maintainer_id,notnull,type:uuid" json:"maintainer_id"`
	Version       uint64    `pg:"version,default:0" json:"version"`
}

func (d *Dataset) Init() *Dataset {
	d.Id = helper.Id()
	if d.ManagerId == "" {
		d.ManagerId = d.MaintainerId
	}
	if d.Visibility == "" {
		d.Visibility = "internal"
	}
	if d.Status == "" {
		d.Status = "inactive"
	}
	return d
}

func (d *Dataset) BasicInfo() iris.Map {
	return iris.Map{
		"id":          d.Id,
		"name":        d.Name,
		"type":        d.Type,
		"message":     d.Message,
		"description": d.Description,
		"visibility":  d.Visibility,
		"status":      d.Status,
		"version":     d.Version,
	}
}
