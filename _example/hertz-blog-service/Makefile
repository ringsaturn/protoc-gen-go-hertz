install:
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/ringsaturn/protoc-gen-go-hertz@latest

.PHONY:gen
gen:
	cd ../.. && go install
	buf build
	buf generate
	protoc-go-inject-tag -input=./api/product/app/v1/v1.pb.go
	go fmt ./...

lint:
	buf lint

format:
	buf format -w
