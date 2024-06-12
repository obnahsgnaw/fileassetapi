
.PHONY: help
help:base_help server_help

.PHONY: base_help
base_help:
	@echo "usage: make <option> <params>"
	@echo "options and effects:"
	@echo "    help   : Show help"

.PHONY: server_help
server_help:
	@echo "server options :"
	@echo "    api         		   : Generate frontend and backend api"
	@echo "    frontend-api		   : Generate frontend api"
	@echo "    backend-api 		   : Generate backend api"
	@echo "    asset       		   : Generate asset package"

.PHONY: init
init:
	@go env -w GOPRIVATE=github.com/obnahsgnaw
	@git config --global url."https://obnahsgnaw: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@github.com".insteadOf "https://github.com"
	@go mod tidy
	@go install github.com/xiaoqidun/gitcz@v1.0.4
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	@go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	@go install github.com/go-bindata/go-bindata/...@v3.1.2


.PHONY: version_help
version_help:
	@echo "version options:"
	@echo "    major    : Generate major version number"
	@echo "    minor    : Generate minor version number"
	@echo "    version  : Generate auto version number"
	@echo "    changelog: Generate change log file and modify tag"

.PHONY: major
major:
	@echo "Version from ${shell git describe --tags `git rev-list --tags --max-count=1`} to v${shell gsemver bump major}"
	@git tag -a "v${shell gsemver bump major}"

.PHONY: minor
minor:
	@echo "Version from ${shell git describe --tags `git rev-list --tags --max-count=1`} to v${shell gsemver bump minor}"
	@git tag -a "v${shell gsemver bump minor}"

.PHONY: version
version:
	@echo "Version from ${shell git describe --tags `git rev-list --tags --max-count=1`} to v${shell gsemver bump pitch}"
	@git tag -a "v${shell gsemver bump}"

.PHONY: changelog
changelog:
	@echo "Generating change log and tag..."
	@git-chglog -o CHANGELOG.md
	@git add CHANGELOG.md
	@git commit -m "chore(release): ${shell git describe --tags `git rev-list --tags --max-count=1`}"
	@git tag -a -f ${shell git describe --tags `git rev-list --tags --max-count=1`}
	@echo "Done"

.PHONY: backend-api
backend-api:
	@echo "Generate backend pb file..."
	@cd backend && buf lint apis --error-format=json && buf generate apis --template buf.gen.yaml
	@protoc -I ./backend/apis \
    -I ./backend/vendors \
    --openapiv2_out=logtostderr=true,allow_merge=true,disable_default_errors=true,merge_file_name=backend:./doc \
    ./backend/apis/*/*/v*/*.proto
	@echo "Done"

.PHONY: frontend-api
frontend-api:
	@echo "Generate frontend pb file..."
	@cd frontend && buf lint apis --error-format=json && buf generate apis --template buf.gen.yaml
	@protoc -I ./frontend/apis \
    -I ./frontend/vendors \
    --openapiv2_out=logtostderr=true,allow_merge=true,disable_default_errors=true,merge_file_name=frontend:./doc \
    ./frontend/apis/*/*/v*/*.proto
	@echo "Done"

.PHONY: asset
asset:
	@echo "Generate document asset file..."
	@go-bindata -o=asset/asset.go -pkg=asset doc/backend.swagger.json doc/frontend.swagger.json #doc/tcp.html doc/wss.html
	@echo "Done"

.PHONY: api
api: backend-api frontend-api asset

.PHONY: socket
socket:tcp-socket wss-socket queue-socket asset

.PHONY: all
all: backend-api frontend-api queue-socket asset

.PHONY: tcp-socket
tcp-socket:
	@echo "Generate frontend tcp socket file..."
	@cd tcp && buf lint apis --error-format=json && buf generate apis --template buf.gen.yaml
	@protoc -I ./tcp/apis \
    -I ./tcp/vendors \
    --doc_out=./doc --doc_opt=html,tcp.html \
    ./tcp/apis/*/*/v*/*.proto
	@echo "Done"
.PHONY: wss-socket
wss-socket:
	@echo "Generate frontend wss socket file..."
	@cd websocket && buf lint apis --error-format=json && buf generate apis --template buf.gen.yaml
	@protoc -I ./websocket/apis \
    -I ./websocket/vendors \
    --doc_out=./doc --doc_opt=html,wss.html \
    ./websocket/apis/*/*/v*/*.proto
	@echo "Done"
.PHONY: queue-socket
queue-socket:
	@echo "Generate frontend queue socket file..."
	@cd queue && buf lint apis --error-format=json && buf generate apis --template buf.gen.yaml
	@echo "Done"

.PHONY: package
package:
	@go run cmd/gen.go