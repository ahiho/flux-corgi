package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"fluxcorgi/cmd/codegen/cmdutils"
)

const (
	svcImplDir = "internal/services"

	implTemplate = `package {{ PACKAGE_NAME }}

import (
	"{{ MODULE_NAME }}/pkg/datastore"

	{{ PKG_IMPORT }}
)

type serviceImpl struct {
	{{ UNIMPLEMENT_SERVER }}
}

func New{{ SERVICE_NAME }}() {{ SEVRICE_SERVER }} {
	return &serviceImpl{}
}
`
	funcTemplate = `package {{ PACKAGE_NAME }}

import (
	"context"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	{{ PKG_IMPORT }}
)
{{ COMMENT }}
func (s *serviceImpl) {{ FN_DECRLARE }} {
	{{ FN_BODY }}
}
`
)

var (
	moduleNamedImplTemplate string
)

func main() {
	args := os.Args[1:]
	pkgDir := "pkg/proto"

	var ignores []string
	if len(args) > 0 {
		pkgDir = args[0]
	}

	if len(args) > 1 && strings.HasPrefix(args[1], "--ignores=") {
		ignores = strings.Split(args[1][len("--ignores="):], ",")
	}

	moduleName, err := cmdutils.GetModuleName()
	if err != nil {
		log.Fatalln(err)
	}

	moduleNamedImplTemplate = strings.ReplaceAll(implTemplate, "{{ MODULE_NAME }}", moduleName)

	pkgFiles, err := cmdutils.LookupDir(pkgDir)
	if err != nil {
		log.Fatalln(err)
	}

	newImports := []string{}
	newCodes := []string{}
	newPkgImports := []string{}

	for _, pkgFile := range pkgFiles {
		if cmdutils.IsStringInArray(ignores, pkgFile.Path[:len(pkgFile.Path)-len("_grpc.pb.go")]) {
			continue
		}
		ni, nc, pg := generateCodes(pkgFile, moduleName)
		newImports = append(newImports, ni...)
		newCodes = append(newCodes, nc...)
		newPkgImports = append(newPkgImports, pg...)
	}

	fmt.Println("\nSERVICE CODE GENERATED COMPLETED!")
	fmt.Println("Please do implement in internal/services")
	if len(newImports) == 0 {
		return
	}
	fmt.Print("\n\033[1;33mNew services generated, you need to edit file internal/server/server.go to register them with 2 steps below\033[0m\n")
	fmt.Print("\n\033[1;36m1. Add imports: \033[0m\n\n")
	fmt.Print("\n\033[1;90m// pkg/proto imports\033[0m\n\n")
	for _, pkgImport := range newPkgImports {
		fmt.Printf("\033[1;32m%s\033[0m\n", pkgImport)
	}
	fmt.Print("\n\033[1;90m// internal/services imports\033[0m\n\n")
	for _, newImport := range newImports {
		fmt.Printf("\033[1;32m%s\033[0m\n", newImport)
	}
	fmt.Print("\n\033[1;36m2. Add below codes to function CompleteServer: \033[0m\n\n")
	fmt.Print("\n\033[1;90mfunc (sa *serverApp) CompleteServer() error {\n...\n\033[0m\n\n")
	for _, newCode := range newCodes {
		// fmt.Println(newCode)
		fmt.Printf("\033[1;32m%s\033[0m\n", newCode)
	}
	fmt.Print("\n\033[1;90m...\n}\033[0m\n\n")
}

func generateCodes(pkgFile cmdutils.PkgFileInfo, moduleName string) ([]string, []string, []string) {
	pbb, err := os.ReadFile(pkgFile.Path)
	if err != nil {
		log.Fatalln(err)
	}
	fileLines := strings.Split(string(pbb), "\n")
	pkgName := ""
	idx0 := 0
	idx1 := 0
	idx2 := 0
	for idx, line := range fileLines {
		if strings.HasPrefix(line, "package ") {
			pkgName = line[len("package "):]
		}

		if strings.HasPrefix(line, "type ") && strings.HasSuffix(line, "Server interface {") {
			idx0 = idx
			idx1 = idx + 1
		}
		if strings.HasPrefix(line, "\tmustEmbedUnimplemented") && strings.HasSuffix(line, "Server()") {
			idx2 = idx
			break
		}
	}

	newImports := []string{}
	newCodes := []string{}
	newPkgImports := []string{}

	// check and generate impl file
	svcPkg := strings.Split(pkgFile.Dir, "/")[2]
	svcServerName := strings.Split(fileLines[idx0], " ")[1]
	serviceName := svcServerName[:len(svcServerName)-len("Server")]
	svcDir := fmt.Sprintf("%v/%v", svcImplDir, svcPkg)

	_, err = os.Stat(svcDir)
	if os.IsNotExist(err) {
		os.Mkdir(svcDir, os.ModePerm)
	}

	implPath := fmt.Sprintf("%v/impl.go", svcDir)
	pkgImport := fmt.Sprintf(`%v "%v/%v"`, pkgName, moduleName, pkgFile.Dir)
	_, err = os.Stat(implPath)
	if os.IsNotExist(err) {
		fmt.Printf("Generate impl for service %v at path %v\n", svcServerName, implPath)
		implContent := strings.ReplaceAll(moduleNamedImplTemplate, "{{ PACKAGE_NAME }}", svcPkg)
		implContent = strings.ReplaceAll(implContent, "{{ PKG_IMPORT }}", pkgImport)
		implContent = strings.ReplaceAll(implContent, "{{ SERVICE_NAME }}", serviceName)
		implContent = strings.ReplaceAll(implContent, "{{ SEVRICE_SERVER }}", fmt.Sprintf("%v.%v", pkgName, svcServerName))
		implContent = strings.ReplaceAll(implContent, "{{ UNIMPLEMENT_SERVER }}", fmt.Sprintf("%v.Unimplemented%v", pkgName, svcServerName))
		implContent = strings.ReplaceAll(implContent, "{{ LOG_SERVICE_NAME }}", cmdutils.ToSnakeCase(pkgName))
		os.WriteFile(implPath, []byte(implContent), 0644)
		exec.Command("goimports", "-w", implPath).Run()

		newImports = append(newImports, fmt.Sprintf(`"%v/%v"`, moduleName, svcDir))
		newPkgImports = append(newPkgImports, pkgImport)
		serviceInstanceName := strings.ToLower(serviceName[0:1]) + serviceName[1:]
		newCodes = append(newCodes, fmt.Sprintf("%v := %v.New%v()", serviceInstanceName, svcPkg, serviceName))
		newCodes = append(newCodes, fmt.Sprintf("sa.RegisterGrpcService(&%v.%v_ServiceDesc, %v)", pkgName, serviceName, serviceInstanceName))
		// // only register handler if gateway existed
		// protoGatewayPath := pkgFile.Path[:len(pkgFile.Path)-len("_grpc.pb.go")] + ".pb.gw.go"
		// _, err = os.Stat(protoGatewayPath)
		// if os.IsNotExist(err) {
		// 	newCodes = append(newCodes, fmt.Sprintf("// %v.Register%vHandler(ctx, sa.GatewayServeMux(), sa.GatewayClientConn())", pkgName, serviceName))
		// } else {
		// 	newCodes = append(newCodes, fmt.Sprintf("%v.Register%vHandler(ctx, sa.GatewayServeMux(), sa.GatewayClientConn())", pkgName, serviceName))
		// }
		// newCodes = append(newCodes, fmt.Sprintf("sa.clients.Register%vClient(%v.New%vClient(sa.InternalClientConn()))", serviceName, pkgName, serviceName))
		// newCodes = append(newCodes, "")
	} else if err != nil {
		log.Fatalln(err)
	}

	// generate files for functions
	functions := fileLines[idx1:idx2]
	comments := []string{}
	for _, fnc := range functions {
		fnName, fnDeclare, fnBody, ok := getFuncInfo(fnc, pkgName)
		if !ok {
			if fnName != "" {
				comments = append(comments, fnName)
			}
			continue
		}
		fnFile := cmdutils.ToSnakeCase(fnName)
		fnPath := fmt.Sprintf("%v/%v.go", svcDir, fnFile)
		_, err = os.Stat(fnPath)
		if os.IsNotExist(err) {
			fmt.Printf("Generate file for function %v at path %v\n", fnDeclare, fnPath)
			funcContent := strings.ReplaceAll(funcTemplate, "{{ PACKAGE_NAME }}", svcPkg)
			funcContent = strings.ReplaceAll(funcContent, "{{ PKG_IMPORT }}", pkgImport)
			funcContent = strings.ReplaceAll(funcContent, "{{ FN_DECRLARE }}", fnDeclare)
			funcContent = strings.ReplaceAll(funcContent, "{{ FN_BODY }}", fnBody)
			if len(comments) > 0 {
				cmt := strings.Join(comments, "\n")
				funcContent = strings.ReplaceAll(funcContent, "{{ COMMENT }}", cmt)
			} else {
				funcContent = strings.ReplaceAll(funcContent, "{{ COMMENT }}", "")
			}
			os.WriteFile(fnPath, []byte(funcContent), 0644)
			exec.Command("goimports", "-w", fnPath).Run()
		} else if err != nil {
			log.Fatalln(err)
		}
	}

	return newImports, newCodes, newPkgImports
}

func getFuncInfo(fnc string, pkgName string) (string, string, string, bool) {
	fnDeclare := strings.Trim(fnc, "\t")
	if fnDeclare == "" {
		return "", "", "", false
	}
	if strings.HasPrefix(fnDeclare, "//") {
		return fnDeclare, "", "", false
	}
	idx01 := strings.Index(fnDeclare, "(")

	fnName := fnDeclare[0:idx01]
	fnInOut := fnDeclare[idx01:]
	var fnIn, fnOut, fnBody string

	if strings.Contains(fnInOut, ") (") {
		// unary
		arr := strings.Split(fnInOut, ") (")
		fnIn = arr[0][1:]
		fnOut = arr[1][:len(arr[1])-1]
		inParams := strings.Split(fnIn, ", ")
		inParamType := inParams[1]
		if !strings.Contains(inParamType, ".") {
			inParamType = fmt.Sprintf("*%v.%v", pkgName, inParamType[1:])
		}
		fnIn = fmt.Sprintf("(ctx context.Context, req %v)", inParamType)

		outParams := strings.Split(fnOut, ", ")
		outParamType := outParams[0]
		if !strings.Contains(outParamType, ".") {
			outParamType = fmt.Sprintf("*%v.%v", pkgName, outParamType[1:])
		}
		fnOut = fmt.Sprintf("(res %v, err error)", outParamType)

		fnDeclare = fmt.Sprintf("%v%v %v", fnName, fnIn, fnOut)
		fnBody = fmt.Sprintf(`return nil, status.Errorf(codes.Unimplemented, "method %v not implemented")`, fnName)
	} else {
		// stream
		arr := strings.Split(fnInOut, ") ")
		fnIn = arr[0][1:]
		if strings.Contains(fnIn, ", ") {
			inParams := strings.Split(fnIn, ", ")
			inParamType0 := inParams[0]
			if !strings.Contains(inParamType0, ".") {
				if inParamType0[0] == '*' {
					inParamType0 = fmt.Sprintf("*%v.%v", pkgName, inParamType0[1:])
				} else {
					inParamType0 = fmt.Sprintf("%v.%v", pkgName, inParamType0)
				}
			}

			inParamType1 := fmt.Sprintf("%v.%v", pkgName, inParams[1])

			fnDeclare = fmt.Sprintf("%v(req %v, sv %v) (err error)", fnName, inParamType0, inParamType1)
		} else {
			fnIn = fmt.Sprintf("%v.%v", pkgName, fnIn)
			fnDeclare = fmt.Sprintf("%v(sv %v) (err error)", fnName, fnIn)
		}

		fnBody = fmt.Sprintf(`return status.Errorf(codes.Unimplemented, "method %v not implemented")`, fnName)
	}

	return fnName, fnDeclare, fnBody, true
}
