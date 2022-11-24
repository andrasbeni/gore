.DEFAULT_GOAL := all

clean:
	@echo Cleaning...
	rm -rf generated/*
	go clean

generate:
	@echo Generating sources...
	mkdir -p generated
	go generate
	protoc --go_out=generated --go_opt=paths=source_relative     --go-grpc_out=generated --go-grpc_opt=paths=source_relative comms/comms.proto

compile: generate
	@echo Compiling...
	go build -o bin/gore

all: clean compile
