install:
	go install github.com/favadi/protoc-go-inject-tag@latest

gen_example:
	buf build
	buf generate
	protoc-go-inject-tag -input=./api/product/app/v1/v1.pb.go -remove_tag_comment
