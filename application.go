package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Application struct {
	tableName struct{} `pg:"alias:application,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Description  string    `pg:"description" json:"description"`
	Token        string    `pg:"token,notnull" json:"token"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
	Tags         []string  `pg:"tags,type:uuid[]" json:"tags"`
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

func (a *Application) Init() *Application {
	id := helper.Id()
	key := helper.AccessKey()
	a.Id = id
	a.Token = helper.AccessId(id, key) + "." + key
	if a.ManagerId == "" {
		a.ManagerId = a.MaintainerId
	}
	if a.Visibility == "" {
		a.Visibility = "internal"
	}
	if a.Status == "" {
		a.Status = "active"
	}
	return a
}

func (a *Application) BasicInfo() iris.Map {
	return iris.Map{
		"id":          a.Id,
		"name":        a.Name,
		"description": a.Description,
		"visibility":  a.Visibility,
		"status":      a.Status,
		"version":     a.Version,
	}
}
