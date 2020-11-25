# PROJECTNAME=$(shell basename "$(PWD)")
PROJECTNAME=tools

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

install:
	go mod download

build:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit

production:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go

start:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit
	./bin/$(PROJECTNAME)

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit
	docker build -t registry.cn-hangzhou.aliyuncs.com/metro/batch_send_coupon:latest .
	docker push registry.cn-hangzhou.aliyuncs.com/metro/batch_send_coupon