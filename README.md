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

### 目录说明

├── CHANGELOG.md  
├── README.md  
├── asset  静态资源打包后  
├── backend  后台接口  
├── doc  生成的文档资源  
├── frontend  前台接口  
├── gen  自动生成目录  
├── go.mod  
├── go.sum  
├── makefile  
├── service  服务  
├── tcp  tcp接口  
├── version  
└── websocket  websocket 接口


### 使用步骤
1. 初始化控的git项目 `git init`
2. 设置远程 frameworkapi 仓库地址 `git remote add github git@github.com:obnahsgnaw/frameworkapi.git`
3. 拉取 framework 分支 `git fetch github framework:framework`
4. 创建本地main分支 `git checkout -b main framework` 如果已存在可以 rebase framework 分支 `git rebase framework`
5. 替换项目包名，全局搜索 `github.com/obnahsgnaw/frameworkapi` 替换为自己的项目的包名，⚠️别修改本README.md的内容
6. 修改 backend/apis/framework_backend_api, 更改目录名称为自己项目关系的名称 xxx_backend_api, 并调整里面proto的packange和go_package, proto文件搜索 framework_ 替换为 xxx_
7. 同理还有 `frontend` `tcp` `websocket`
8. 删除gen目录下内容，`make all` 重新生成文件，并全局替换 `framework_` 为 自定义的 `xxx_`
9. 添加自定义的相关proto