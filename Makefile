buf_generate:
	buf generate
	go run cmd/codegen/swaggerutils/main.go swagger/docs.swagger.json
	rm -rf tmp
	mkdir tmp
	openapi-generator-cli generate
	jq 'del(.servers)' tmp/openapi.json > swagger/docs.swagger.json
	rm -rf tmp
	@if [ "$(shell uname)" = "Darwin" ]; then\
		sed -i '' 's/"scheme" : "basic",/"scheme" : "bearer",/g' swagger/docs.swagger.json;\
		sed -i '' 's/"scheme": "basic",/"scheme": "bearer",/g' swagger/docs.swagger.json;\
	else\
		sed -i 's/"scheme" : "basic",/"scheme" : "bearer",/g' swagger/docs.swagger.json;\
		sed -i 's/"scheme":"basic",/"scheme":"bearer",/g' swagger/docs.swagger.json;\
		sed -i 's/"scheme": "basic",/"scheme": "bearer",/g' swagger/docs.swagger.json;\
	fi
	go run cmd/codegen/internalsvc/main.go pkg/proto
	go run cmd/codegen/gateway/main.go
	gofmt -w pkg/server/proto_generated.go

run:
	ENV=DEVELOPMENT go run cmd/app/main.go