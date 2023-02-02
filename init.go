package iris_common_model

import (
	server "github.com/AlexZ33/iris-extend-server"
	"github.com/kataras/iris/v12"
	"time"
)

var (
	Schema  iris.Map
	Citus   iris.Map
	Builtin iris.Map
	Indexes map[string][]string
)

func init() {
	current := time.Now()
	Schema = iris.Map{
		"Resource":    new(Resource),
		"Component":   new(Component),
		"Repository":  new(Repository),
		"Collection":  new(Collection),
		"Field":       new(Field),
		"Tag":         new(Tag),
		"Dataset":     new(Dataset),
		"Record":      new(Record),
		"Log":         new(Log),
		"Application": new(Application),
		"Order":       new(Order),
		"Project":     new(Project),
		"Form":        new(Form),
	}
	Citus = iris.Map{}
	Indexes = map[string][]string{
		"Application": {
			"btree(created_at DESC)",
		},
		"Order": {
			"btree(created_at DESC)",
			"hash(application_id)",
			"hash(dataset_id)",
		},
		"Repository": {
			"btree(created_at DESC)",
			"hash(status)",
			"gin(tenants)",
		},
		"Collection": {
			"btree(created_at DESC)",
			"hash(repository_id)",
			"hash(customer_id)",
		},
		"Field": {
			"btree(created_at DESC)",
			"hash(collection_id)",
		},
		"Tag": {
			"btree(created_at DESC)",
		},
		"Resource": {
			"btree(created_at DESC)",
		},
		"Component": {
			"btree(created_at DESC)",
			"gin(tenants)",
		},
		"Project": {
			"btree(created_at DESC)",
			"gin(tenants)",
		},
		"Dataset": {
			"btree(created_at DESC)",
			"hash(project_id)",
			"hash(customer_id)",
		},
		"Record": {
			"btree(recorded_at DESC)",
			"gin(content)",
		},
		"Log": {
			"btree(recorded_at DESC)",
			"gin(content)",
		},
		"Form": {
			"btree(created_at DESC)",
			"hash(dataset_id)",
			"hash(customer_id)",
		},
	}
	Builtin = iris.Map{}
	if server.ProvidesAuth {
		adminUser := (&User{
			Name:       "Admin",
			Account:    "admin",
			Password:   "admin",
			Role:       "admin",
			Visibility: "private",
			CreatedAt:  current,
			UpdatedAt:  current,
		}).Init(true)
		adminUserId := adminUser.Id
		adminUser.ManagerId = adminUserId
		adminUser.MaintainerId = adminUserId
		workerUser := (&User{
			Name:         "Worker",
			Account:      "worker",
			Password:     "worker",
			Role:         "worker",
			Visibility:   "private",
			CreatedAt:    current,
			UpdatedAt:    current,
			MaintainerId: adminUserId,
		}).Init(true)
		workerUserId := workerUser.Id
		auditorUser := (&User{
			Name:         "Auditor",
			Account:      "auditor",
			Password:     "auditor",
			Role:         "auditor",
			Visibility:   "private",
			CreatedAt:    current,
			UpdatedAt:    current,
			MaintainerId: adminUserId,
		}).Init(true)
		if server.MaintainerId != "" {
			auditorUser.Id = server.MaintainerId
		} else {
			server.MaintainerId = auditorUser.Id
		}
		auditorUserId := auditorUser.Id
		rootGroup := (&Group{
			Name:         "Root",
			Members:      []string{adminUserId, workerUserId, auditorUserId},
			Visibility:   "private",
			CreatedAt:    current,
			UpdatedAt:    current,
			MaintainerId: workerUserId,
		}).Init()
		rootGroupId := rootGroup.Id
		superPolicy := (&Policy{
			Name:         "Super",
			TenantId:     rootGroupId,
			ValidFrom:    current,
			ExpiresAt:    current.AddDate(100, 0, 0),
			Visibility:   "private",
			CreatedAt:    current,
			UpdatedAt:    current,
			MaintainerId: auditorUserId,
		}).Init()
		Builtin["AdminUser"] = adminUser
		Builtin["WorkerUser"] = workerUser
		Builtin["AuditorUser"] = auditorUser
		Builtin["RootGroup"] = rootGroup
		Builtin["SuperPolicy"] = superPolicy
		Schema["User"] = new(User)
		Schema["Group"] = new(Group)
		Schema["Policy"] = new(Policy)
		Indexes["User"] = []string{
			"btree(created_at DESC)",
		}
		Indexes["Group"] = []string{
			"btree(created_at DESC)",
			"gin(members)",
		}
		Indexes["Policy"] = []string{
			"btree(created_at DESC)",
		}
	}
}
