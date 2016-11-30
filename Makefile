build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/consul-kv-linux
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o release/consul-kv-osx
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/consul-kv-windows
