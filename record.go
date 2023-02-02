package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Record struct {
	tableName struct{} `pg:"alias:record,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Collection   string    `pg:"collection,notnull" json:"collection"`
	Content      iris.Map  `pg:"content,notnull,type:jsonb" json:"content"`
	RecordedAt   time.Time `pg:"recorded_at,notnull" json:"recorded_at"`
	Integrity    string    `pg:"integrity,notnull" json:"integrity"`
	Signature    string    `pg:"signature,notnull" json:"signature"`
	Visibility   string    `pg:"visibility,default:'internal'" json:"visibility"`
	Status       string    `pg:"status,default:'active'" json:"status"`
	Metrics      iris.Map  `pg:"metrics,type:jsonb" json:"metrics"`
	OtherInfo    iris.Map  `pg:"other_info,type:jsonb" json:"other_info"`
	CreatedAt    time.Time `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt    time.Time `pg:"updated_at,default:now()" json:"updated_at"`
	MaintainerId string    `pg:"maintainer_id,notnull,type:uuid" json:"maintainer_id"`
	Version      uint64    `pg:"version,default:0" json:"version"`
}

func (r *Record) Init() *Record {
	r.Id = helper.Id()
	if r.Visibility == "" {
		r.Visibility = "internal"
	}
	if r.Status == "" {
		r.Status = "active"
	}
	return r
}

func (r *Record) BasicInfo() iris.Map {
	return iris.Map{
		"id":          r.Id,
		"collection":  r.Collection,
		"integrity":   r.Integrity,
		"visiibility": r.Visibility,
		"status":      r.Status,
		"version":     r.Version,
	}
}
