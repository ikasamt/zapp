mkdir_bin:
	mkdir -p ./bin

build: mkdir_bin
	GIN_MODE=release GO111MODULE=on go build -o bin/server

debug: go_generate build
	bin/server

linux: go_generate mkdir_bin 
	GIN_MODE=release GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o bin/server

go_generate: clean
	go generate ./...

download:
	GO111MODULE=on go mod download

clean:
	rm -rf zzz-autogen-*
