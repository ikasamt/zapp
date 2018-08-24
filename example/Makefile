mkdir_bin:
	mkdir -p ./bin

build: mkdir_bin
	GIN_MODE=release go build -o bin/server
	rm -rf zzz-autogen-*

debug: go_generate build
	export TEAME_CONFIG_FILENAME='./configs/local.toml' && bin/server
	rm -rf zzz-autogen-*

linux: go_generate mkdir_bin 
	GIN_MODE=release GOOS=linux GOARCH=amd64 go build -o bin/server
	rm -rf zzz-autogen-*

go_generate: clean
	go generate ./...

clean:
	rm -rf zzz-autogen-*
