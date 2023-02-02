package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Screen struct {
	tableName struct{} `pg:"alias:screen,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Description  string    `pg:"description" json:"description"`
	CustomerId   string    `pg:"customer_id,type:uuid" json:"customer_id"`
	ValidFrom    time.Time `pg:"valid_from,default:now()" json:"valid_from"`
	ExpiresAt    time.Time `pg:"expires_at,default:now()" json:"expires_at"`
	Content      iris.Map  `pg:"content,type:jsonb" json:"content"`
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

func (s *Screen) Init() *Screen {
	s.Id = helper.Id()
	if s.ManagerId == "" {
		s.ManagerId = s.MaintainerId
	}
	if s.Visibility == "" {
		s.Visibility = "internal"
	}
	if s.Status == "" {
		s.Status = "active"
	}
	return s
}

func (s *Screen) BasicInfo() iris.Map {
	return iris.Map{
		"id":          s.Id,
		"name":        s.Name,
		"description": s.Description,
		"visibility":  s.Visibility,
		"status":      s.Status,
		"version":     s.Version,
	}
}
