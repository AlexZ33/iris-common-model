package iris_common_model

import (
	"github.com/kataras/iris/v12"
	"time"
)

var (
	Schema  iris.Map
	Citus   iris.Map
	Builtin iris.Map
	Indexs  map[string][]string
)

func init() {
	current := time.Now()
	Schema = iris.Map{
		"Record": new(Record),
	}
}
