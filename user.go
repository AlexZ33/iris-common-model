package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"github.com/pelletier/go-toml"
	"net"
	"strings"
	"time"
)

type User struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Account      string    `json: "account"`
	Password     string    `json:"password"`
	Mobile       string    `json:"mobile"`
	Email        string    `json:"email"`
	Avatar       string    `json:"avatar"`
	Description  string    `json:"description"`
	LoggedIP     net.IP    `json: "logged_ip""`
	LoggedIn     time.Time `json:"logged-in"`
	LoggedOut    time.Time `json:"logged_out"`
	Role         string    `json: "role"`
	ManagerId    string    `json: "manager_id"`
	Visibility   string    `json: "visivility"`
	Status       string    `json: "status"`
	Metrics      iris.Map  `json: "metrics"`
	BaseInfo     iris.Map  `json: "base_info"`
	OtherInfo    iris.Map  `json: "other_info"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json: "update_at"`
	MaintainerId string    `json: "maintainer_id"`
	Version      uint64    `json: "version"`
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
