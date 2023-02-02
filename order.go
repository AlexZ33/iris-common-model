package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Order struct {
	tableName struct{} `pg:"alias:order,discard_unknown_columns"`

	Id            string    `pg:"id,pk,type:uuid" json:"id"`
	Name          string    `pg:"name,notnull" json:"name"`
	Description   string    `pg:"description" json:"description"`
	ApplicationId string    `pg:"application_id,notnull,type:uuid" json:"application_id"`
	DatasetId     string    `pg:"dataset_id,notnull,type:uuid" json:"dataset_id"`
	Content       iris.Map  `pg:"content,type:jsonb" json:"content"`
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

func (o *Order) Init() *Order {
	o.Id = helper.Id()
	if o.ManagerId == "" {
		o.ManagerId = o.MaintainerId
	}
	if o.Visibility == "" {
		o.Visibility = "internal"
	}
	if o.Status == "" {
		o.Status = "inactive"
	}
	return o
}

func (o *Order) BasicInfo() iris.Map {
	return iris.Map{
		"id":          o.Id,
		"name":        o.Name,
		"message":     o.Message,
		"description": o.Description,
		"visibility":  o.Visibility,
		"status":      o.Status,
		"version":     o.Version,
	}
}
