# protoc-gen-go 扩展

tkform google.golang.org/protobuf@v1.25.0
tkgrpc github.com/golang/protobuf@v1.4.3

## 安装

go get github.com\ikaiguang\protoc-gen-go\cmd\protoc-gen-go-tkform
go get github.com\ikaiguang\protoc-gen-go\cmd\protoc-gen-go-tkgrpc

## protoc-gen-go-tkgrpc

当前版本：github.com/golang/protobuf@v1.4.3

github.com/ikaiguang/srv_toolkit/cmd/protoc-gen-go-tkgrpc/gengogrpc

gengogrpc 源码复制于

github.com/golang/protobuf@v1.4.3
github.com/golang/protobuf/internal/gengogrpc 

## protoc-gen-go-tkform

基于 google.golang.org/protobuf@v1.25.0

## 编辑修改

复制 google.golang.org/protobuf@v1.25.0 -> github.com/ikaiguang/protoc-gen-go

替换引入包
google.golang.org/protobuf -> github.com/ikaiguang/protoc-gen-go

重新复制 google.golang.org/protobuf/cmd/protoc-gen-go/main.go -> github.com/ikaiguang/protoc-gen-go/cmd/protoc-gen-go/main.go
重新复制 google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo -> github.com/ikaiguang/protoc-gen-go/cmd/protoc-gen-go/internal_gengo

修正 ./cmd/protoc-gen-go-tkform
添加 tag form:

## 格式化代码 & 尝试编译

gofmt -w .
go build ./cmd/protoc-gen-go-tkform/main.go
