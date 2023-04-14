package main

import (
	"fluxcorgi/cmd/codegen/cmdutils"
	"fmt"
	"log"
	"os"
	"strings"
)

const gatewayTemplate = `package server

import (
	"context"
	
${{ PKG_IMPORT }}
)

func (s *incompletedServer) RegisterGateway(ctx context.Context) {
${{ GATEWAY_SETUP }}
}
`

func main() {
	args := os.Args[1:]
	pkgDir := "pkg/proto"
	if len(args) > 0 {
		pkgDir = args[0]
	}

	moduleName, err := cmdutils.GetModuleName()
	if err != nil {
		log.Fatalln(err)
	}

	pkgFiles, err := cmdutils.LookupDir(pkgDir)
	if err != nil {
		log.Fatalln(err)
	}

	pkgImports := []string{}
	gatewayRegisters := []string{}

	for _, pkgFile := range pkgFiles {
		protoGatewayPath := pkgFile.Path[:len(pkgFile.Path)-len("_grpc.pb.go")] + ".pb.gw.go"
		_, err = os.Stat(protoGatewayPath)
		if os.IsNotExist(err) {
			continue
		}

		pbb, err := os.ReadFile(protoGatewayPath)
		if err != nil {
			log.Fatalln(err)
		}
		fileLines := strings.Split(string(pbb), "\n")
		pkgName := ""
		pkgImportPath := ""

		for _, fileLine := range fileLines {
			if strings.HasPrefix(fileLine, "package ") && pkgName == "" {
				pkgName = fileLine[len("package "):]
				pkgImportPath = "\t" + fmt.Sprintf(`%v "%v/%v"`, pkgName, moduleName, pkgFile.Dir)
			}

			if strings.HasPrefix(fileLine, "func Register") && strings.Contains(fileLine, "Handler(ctx") {
				fnName := fileLine[len("func "):strings.Index(fileLine, "(")]
				gatewayRegisters = append(gatewayRegisters, fmt.Sprintf("\t%v.%v(ctx, s.gatewayMux, s.gatewayClientConn)", pkgName, fnName))
			}

		}

		if !cmdutils.IsStringInArray(pkgImports, pkgImportPath) {
			pkgImports = append(pkgImports, pkgImportPath)
		}
	}

	generatedGateway := strings.ReplaceAll(gatewayTemplate, "${{ PKG_IMPORT }}", strings.Join(pkgImports, "\n"))
	generatedGateway = strings.ReplaceAll(generatedGateway, "${{ GATEWAY_SETUP }}", strings.Join(gatewayRegisters, "\n"))

	os.WriteFile("pkg/server/proto_generated.go", []byte(generatedGateway), 0644)
}
