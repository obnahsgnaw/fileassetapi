package frontend

import (
	"github.com/obnahsgnaw/api"
	// TODO 增加新增的proto模块
	_ "github.com/obnahsgnaw/fileassetapi/gen/fileasset_backend_api/index/v1"
	"github.com/obnahsgnaw/fileassetapi/service"
)

func Scan(s *api.Server) error {
	if s == nil {
		return nil
	}
	return service.ScanNoAuth(s, "fileasset_frontend_api")
}
