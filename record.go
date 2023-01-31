package iris_common_model

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	"github.com/kataras/iris/v12"
	"time"
)

type Record struct {
	Id           string    `json:"id"`
	Collection   string    `json:"collection"`
	Content      iris.Map  `json:"content"`
	RecordedAt   time.Time `json:"recorded_at"`
	Integrity    string    `json:"integrity"`
	Signature    string    `json:"signature"`
	Visibility   string    `json:"visibility"`
	Status       string    `json:"status"`
	Metrics      iris.Map  `json:"metrics"`
	OtherInfo    iris.Map  `json:"other_info"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MaintainerId string    `json:"maintainer_id"`
	Version      uint64    `json:"version"`
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
