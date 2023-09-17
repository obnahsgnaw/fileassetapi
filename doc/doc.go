package doc

import "github.com/obnahsgnaw/frameworkapi/asset"

func FrontendDoc() ([]byte, error) {
	return asset.Asset("doc/frontend.swagger.json")
}

func BackendDoc() ([]byte, error) {
	return asset.Asset("doc/backend.swagger.json")
}

func TcpDoc() ([]byte, error) {
	return asset.Asset("doc/tcp.html")
}

func WssDoc() ([]byte, error) {
	return asset.Asset("doc/wss.html")
}
