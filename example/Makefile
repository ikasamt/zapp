mkdir_bin:
	mkdir -p ./bin

build: mkdir_bin
	GIN_MODE=release go build -o bin/server

debug: build
	export TEAME_CONFIG_FILENAME='./configs/local.toml' && bin/server

linux: mkdir_bin 
	GIN_MODE=release GOOS=linux GOARCH=amd64 go build -o bin/server

