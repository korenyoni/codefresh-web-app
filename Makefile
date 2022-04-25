build_path = build
name = webapp

.PHONY: all build

clean:
	rm -rf $(build_path)

run:
	go run main.go

build:
	GOOS=linux GOARCH=arm64 go build -o $(build_path)/$(name)-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o $(build_path)/$(name)-linux-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o $(build_path)/$(name)-darwin-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o $(build_path)/$(name)-darwin-amd64 main.go

all: clean build