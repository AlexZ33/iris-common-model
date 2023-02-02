package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"github.com/pelletier/go-toml"
	"net"
	"strings"
	"time"
)

type User struct {
	tableName struct{} `pg:"alias:user,discard_unknown_columns"`

	Id           string    `pg:"id,pk,type:uuid" json:"id"`
	Name         string    `pg:"name,notnull" json:"name"`
	Account      string    `pg:"account,notnull" json:"account"`
	Password     string    `pg:"password,notnull" json:"password"`
	Mobile       string    `pg:"mobile" json:"mobile"`
	Email        string    `pg:"email" json:"email"`
	Avatar       string    `pg:"avatar" json:"avatar"`
	Description  string    `pg:"description" json:"description"`
	LoggedIP     net.IP    `pg:"logged_ip" json:"logged_ip"`
	LoggedIn     time.Time `pg:"logged_in,default:now()" json:"logged_in"`
	LoggedOut    time.Time `pg:"logged_out,default:now()" json:"logged_out"`
	Role         string    `pg:"role,default:'user'" json:"role"`
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

// 初始化
func (u *User) Init(raw bool, conf *toml.Tree) *User {
	if raw {
		config := GetTree(conf, "crypto")
		password := Hash(u.Password, config, false)
		token := Token(u.Account, password)
		u.Password = Hash(token, config, true)
	}
	u.Id = Id()
	u.Metrics = iris.Map{
		"login_count":         0,
		"fail_login_attempts": 0,
	}
	if u.ManagerId == "" {
		u.ManagerId = u.MaintainerId
	}
	if u.Visibility == "" {
		u.Visibility = "internal"
	}
	if u.Status == "" {
		u.Status = "active"
	}
	return u
}

// 用户基本信息
func (u *User) BasicInfo() iris.Map {
	return iris.Map{
		"id":          u.Id,
		"name":        u.Name,
		"description": u.Description,
		"role":        u.Role,
		"visibility":  u.Visibility,
		"status":      u.Status,
		"version":     u.Version,
	}
}

func (u *User) GetMap() iris.Map {
	return iris.Map{
		"id":            u.Id,
		"name":          u.Name,
		"mobile":        u.Mobile,
		"email":         u.Email,
		"avatar":        u.Avatar,
		"description":   u.Description,
		"role":          u.Role,
		"manager_id":    u.ManagerId,
		"visibility":    u.Visibility,
		"status":        u.Status,
		"other_info":    u.OtherInfo,
		"logged_ip":     u.LoggedIP,
		"logged_in":     u.LoggedIn,
		"logged_out":    u.LoggedOut,
		"created_at":    u.CreatedAt,
		"updated_at":    u.UpdatedAt,
		"maintainer_id": u.MaintainerId,
		"version":       u.Version,
	}

}

func (u *User) CheckRole(prefix string) bool {
	for _, role := range strings.Split(u.Role, ",") {
		if strings.HasPrefix(role, prefix) || strings.HasPrefix(prefix, role) {
			return true
		}
	}
	return false
}
