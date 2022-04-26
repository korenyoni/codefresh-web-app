build_path 	= build
name 		= webapp
github_org 	= korenyoni
github_repo = codefresh-web-app

.PHONY: all build docker-build

clean:
	rm -rf $(build_path)

run:
	PORT=8080 go run main.go

build:
	GOOS=linux GOARCH=arm64 go build -o $(build_path)/$(name)-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o $(build_path)/$(name)-linux-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o $(build_path)/$(name)-darwin-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o $(build_path)/$(name)-darwin-amd64 main.go

docker-build:
	docker build --build-arg GITHUB_ORG=$(github_org) --build-arg GITHUB_REPO=$(github_repo) -t $(github_org)/$(github_repo) .

all: clean build docker-build