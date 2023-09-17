package service

import (
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/obnahsgnaw/api"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"strings"
)

// ScanNoAuth backendapi  frontendapi
func ScanNoAuth(s *api.Server, prefix string) error {
	var err error
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if !strings.HasPrefix(fd.Path(), prefix) {
			return true
		}
		// client ignore
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			service := services.Get(i)
			methods := service.Methods()
			for j := 0; j < methods.Len(); j++ {
				method := methods.Get(j)
				noAuth := true
				methodName := "GET"
				uri := ""
				if ope, _ := proto.GetExtension(method.Options(), options.E_Openapiv2Operation).(*options.Operation); ope != nil {
					securities := ope.Security
					for _, security := range securities {
						if _, secExist := security.SecurityRequirement["BearerToken"]; secExist {
							noAuth = false
							break
						}
					}
				} else {
					continue
				}
				if apiInfo, _ := proto.GetExtension(method.Options(), annotations.E_Http).(*annotations.HttpRule); apiInfo != nil {
					if uri = apiInfo.GetGet(); uri != "" {
						methodName = "GET"
					} else if uri = apiInfo.GetPost(); uri != "" {
						methodName = "POST"
					} else if uri = apiInfo.GetPut(); uri != "" {
						methodName = "PUT"
					} else if uri = apiInfo.GetPatch(); uri != "" {
						methodName = "PATCH"
					} else if uri = apiInfo.GetDelete(); uri != "" {
						methodName = "DELETE"
					} else {
						err = errors.New("resource scan: Unknown method")
						return false
					}
				}
				if uri == "" || methodName == "" {
					err = errors.New("resource scan: no method uri or method type[" + string(method.FullName()) + "]")
					return false
				}
				if noAuth {
					s.AuthRouteManager().AddAuthIgnoredRoute(methodName, uri)
				}
			}
		}
		return true
	})
	return err
}
