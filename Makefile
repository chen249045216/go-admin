all: run

assets:
	go-bindata -o ./template/adminlte/resource/assets.go ./template/adminlte/resource/assets/...

tmpl:
	admincli compile tpl

fmt:
	go fmt ./adapter/...
	go fmt ./admincli/...
	go fmt ./context/...
	go fmt ./engine/...
	go fmt ./examples/...
	go fmt ./modules/...
	go fmt ./plugins/...
	go fmt ./template/...