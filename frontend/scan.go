package frontend

import (
	"github.com/obnahsgnaw/api"
	// TODO 增加新增的proto模块
	_ "github.com/obnahsgnaw/frameworkapi/gen/framework_backend_api/index/v1"
	"github.com/obnahsgnaw/frameworkapi/service"
)

func Scan(s *api.Server) error {
	if s == nil {
		return nil
	}
	return service.ScanNoAuth(s, "framework_frontend_api")
}
