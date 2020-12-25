# protoc-gen-go 扩展

tkform google.golang.org/protobuf@v1.25.0
tkgrpc github.com/golang/protobuf@v1.4.3

## protoc-gen-go-tkgrpc

当前版本：github.com/golang/protobuf@v1.4.3

github.com/ikaiguang/srv_toolkit/cmd/protoc-gen-go-tkgrpc/gengogrpc

gengogrpc 源码复制于

github.com/golang/protobuf@v1.4.3
github.com/golang/protobuf/internal/gengogrpc 

## protoc-gen-go-tkform

基于 google.golang.org/protobuf@v1.25.0

## 替换引入包

google.golang.org/protobuf
github.com/ikaiguang/protoc-gen-go

## 格式化代码 & 尝试编译

gofmt -w .
go build ./cmd/protoc-gen-go-tkform/main.go
