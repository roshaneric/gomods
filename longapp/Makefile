NAME   := longapp
VER    := v0.0.1
GO_VER=$(shell go version|grep "go version"|cut -d' ' -f3|sed "s/[\s\t]*//")

run:
	go run ./cmd

run2:
	rm -rf bin/${NAME}; go build -i -ldflags "-X main.Revision=$(REV) -X main.GoVersion=$(GO_VER)" -o bin/${NAME} ./cmd;bin/${NAME} -conf="config.yaml"
