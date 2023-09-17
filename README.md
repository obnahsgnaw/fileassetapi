##  Api define

### 环境依赖
1. 安装 `protoc@3.21.12`
2. 安装 `buf@1.15.0`
3. 安装 `make` 环境
4. protobuf插件
    - protoc-gen-go:`go install google.golang.org/protobuf/cmd/protoc-gen-go`
    - protoc-gen-go-grpc:`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0`
    - protoc-gen-grpc-gateway:`go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2`
    - protoc-gen-openapiv2:`go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2`
    - protoc-gen-validate: `go install github.com/envoyproxy/protoc-gen-validate@v0.10.1`
    - bin data: `go install github.com/go-bindata/go-bindata/...@v3.1.2`

include `backend api`、`frontend api`、`socket tcp and websocket api`

1. 全局替换包名 “github.com/obnahsgnaw/frameworkapi”
2. 改apis下的目录名及proto包名
3. 从新生成proto