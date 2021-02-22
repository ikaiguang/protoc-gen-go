package tarsrpc

import (
	"fmt"
	"github.com/ikaiguang/protoc-gen-go/compiler/protogen"
	"strings"
)

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	contextPath    = protogen.GoImportPath("context")
	fmtPath        = protogen.GoImportPath("fmt")
	modelPkgPath   = protogen.GoImportPath("github.com/TarsCloud/TarsGo/tars/model")
	requestPkgPath = protogen.GoImportPath("github.com/TarsCloud/TarsGo/tars/protocol/res/requestf")
	tarsPkgPath    = protogen.GoImportPath("github.com/TarsCloud/TarsGo/tars")
	toolsPath      = protogen.GoImportPath("github.com/TarsCloud/TarsGo/tars/util/tools")
	currentPath    = protogen.GoImportPath("github.com/TarsCloud/TarsGo/tars/util/current")
	protoPath      = protogen.GoImportPath("google.golang.org/protobuf/proto")
)

// upperFirstLatter make the fisrt charater of given string  upper class
func upperFirstLatter(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return strings.ToUpper(string(s[0]))
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// GenerateFile generates a _grpc.pb.go file containing gRPC service definitions.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_tars.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-tktars. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	GenerateFileContent(gen, file, g)
	return g
}

// GenerateFileContent generates the gRPC service definitions, excluding the package statement.
func GenerateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	g.QualifiedGoIdent(fmtPath.Ident("fmt"))
	g.QualifiedGoIdent(contextPath.Ident("context"))
	g.QualifiedGoIdent(protoPath.Ident("proto"))
	g.QualifiedGoIdent(modelPkgPath.Ident("model"))
	g.QualifiedGoIdent(requestPkgPath.Ident("requestf"))
	g.QualifiedGoIdent(tarsPkgPath.Ident("tars"))
	g.QualifiedGoIdent(toolsPath.Ident("tools"))
	g.QualifiedGoIdent(currentPath.Ident("current"))
	//g.Import(modelPkgPath)
	//g.Import(requestPkgPath)
	//g.Import(tarsPkgPath)
	//g.Import(toolsPath)
	//g.Import(currentPath)
	//g.Import(contextPath)
	for i, service := range file.Services {
		generateService(g, file, service, i)
	}
}

// generateService .
func generateService(g *protogen.GeneratedFile, file *protogen.File, service *protogen.Service, index int) {
	originServiceName := service.GoName
	serviceName := upperFirstLatter(originServiceName)
	g.P("// This following code was generated by tarsrpc")
	g.P(fmt.Sprintf("// Gernerated from %s", file.Desc.Name()))
	g.P(fmt.Sprintf(`type  %s struct {
		s model.Servant
	}
	`, serviceName))
	g.P()

	//generate SetServant
	g.P(fmt.Sprintf(`//SetServant is required by the servant interface.
	func (obj *%s) SetServant(s model.Servant){
		obj.s = s
	}
	`, serviceName))
	g.P()
	//generate AddServant
	g.P(fmt.Sprintf(`//AddServant is required by the servant interface
	func (obj *%s) AddServant(imp imp%s, objStr string){
		tars.AddServant(obj, imp, objStr)
	}`, serviceName, serviceName))

	//generate AddServantWithContext
	g.P(fmt.Sprintf(`////AddServant adds servant  for the service with context
	func (obj *%s) AddServantWithContext(imp imp%sWithContext, objStr string) {
		tars.AddServantWithContext(obj, imp, objStr)
	}`, serviceName, serviceName))
	g.P()

	//generate TarsSetTimeout
	g.P(fmt.Sprintf(`//TarsSetTimeout is required by the servant interface. t is the timeout in ms. 
	func (obj *%s) TarsSetTimeout(t int){
		obj.s.TarsSetTimeout(t)
	}
	`, serviceName))
	g.P()

	//generate TarsSetProtocol
	g.P(fmt.Sprintf(`//TarsSetProtocol is required by the servant interface. t is the protocol. 
	func (obj *%s) TarsSetProtocol(p model.Protocol){
		obj.s.TarsSetProtocol(p)
	}
	`, serviceName))
	g.P()

	//generate the interface
	g.P(fmt.Sprintf("type imp%s interface{", serviceName))
	for _, method := range service.Methods {
		g.P(fmt.Sprintf("%s (input %s) (output %s, err error)",
			upperFirstLatter(method.GoName), g.QualifiedGoIdent(method.Input.GoIdent), g.QualifiedGoIdent(method.Output.GoIdent)))
	}
	g.P("}")
	g.P()

	//generate the context interface
	g.P(fmt.Sprintf("type imp%sWithContext interface{", serviceName))
	for _, method := range service.Methods {
		g.P(fmt.Sprintf("%s (ctx context.Context, input %s) (output %s, err error)",
			upperFirstLatter(method.GoName), g.QualifiedGoIdent(method.Input.GoIdent), g.QualifiedGoIdent(method.Output.GoIdent)))
	}
	g.P("}")
	g.P()

	//gernerate the dispathcer
	generateDispatch(g, service)

	for _, method := range service.Methods {
		generateClientCode(g, service, method)
	}
}

// generateClientCode .
func generateClientCode(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	methodName := upperFirstLatter(method.GoName)
	serviceName := upperFirstLatter(service.GoName)
	inType := g.QualifiedGoIdent(method.Input.GoIdent)
	outType := g.QualifiedGoIdent(method.Output.GoIdent)
	g.P(fmt.Sprintf(`// %s is client rpc method as defined
		func (obj *%s) %s(input %s, _opt ...map[string]string)(output %s, err error){
			ctx := context.Background()
			return obj.%sWithContext(ctx, input, _opt...)
		}
	`, methodName, serviceName, methodName, inType, outType, methodName))

	g.P(fmt.Sprintf(`// %sWithContext is client rpc method as defined
		func (obj *%s) %sWithContext(ctx context.Context, input %s, _opt ...map[string]string)(output %s, err error){
			var inputMarshal []byte
			inputMarshal, err = proto.Marshal(&input)
			if err != nil {
				return output, err
			}

			var _status map[string]string
			var _context map[string]string
			if len(_opt) == 1 {
				_context = _opt[0]
			} else if len(_opt) == 2 {
				_context = _opt[0]
				_status = _opt[1]
			}

			resp := new(requestf.ResponsePacket)

			err = obj.s.Tars_invoke(ctx, 0, "%s", inputMarshal, _status, _context, resp)
			if err != nil {
				return output, err
			}
			if err = proto.Unmarshal(tools.Int8ToByte(resp.SBuffer), &output); err != nil{
				return output, err
			}

			if len(_opt) == 1 {
				for k := range _context {
					delete(_context, k)
				}
				for k, v := range resp.Context {
					_context[k] = v
				}
			} else if len(_opt) == 2 {
				for k := range _context {
					delete(_context, k)
				}
				for k, v := range resp.Context {
					_context[k] = v
				}
				for k := range _status {
					delete(_status, k)
				}
				for k, v := range resp.Status {
					_status[k] = v
				}
			}

			return output, nil
		}
	`, methodName, serviceName, methodName, inType, outType, method.GoName))
}

// generateDispatch .
func generateDispatch(g *protogen.GeneratedFile, service *protogen.Service) {
	serviceName := upperFirstLatter(service.GoName)
	g.P(fmt.Sprintf(`//Dispatch is used to call the user implement of the defined method.
	func (obj *%s) Dispatch(ctx context.Context, val interface{}, req * requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool)(err error){
		input := tools.Int8ToByte(req.SBuffer)
		var output []byte
		funcName := req.SFuncName
		switch funcName {
	`, serviceName))
	for _, method := range service.Methods {
		g.P(fmt.Sprintf(`case "%s":
			inputDefine := %s{}
			if err = proto.Unmarshal(input,&inputDefine); err != nil{
				return err
			}
			var res %s
            if withContext == false {
				imp := val.(imp%s)
				res, err = imp.%s(inputDefine)
				if err != nil {
					return err
				}
			}else {
				imp := val.(imp%sWithContext)
				res, err = imp.%s(ctx, inputDefine)
				if err != nil {
					return err
				}
			}
			output , err = proto.Marshal(&res)
			if err != nil {
				return err
			}
		`, method.GoName, g.QualifiedGoIdent(method.Input.GoIdent), g.QualifiedGoIdent(method.Output.GoIdent), serviceName, upperFirstLatter(method.GoName), serviceName, upperFirstLatter(method.GoName)))
	}
	g.P(`default:
			return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(output),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
	`)
	g.P()
}
