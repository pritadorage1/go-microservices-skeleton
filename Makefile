GOPATH:=$(shell go env GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
PROTOCMD=protoc

grpc:
	# building grpc binary
	$(GOBUILD) -o bin/grpc cmd/grpc/main.go
rest:
	# building rest gateway binary
	$(GOBUILD) -o bin/rest cmd/rest/main.go
proto:
	# compiling protobuffer code:
	# $(PROTOCMD) --proto_path=api \
	# -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	# --micro_out=./usecase \
	# --go_out=plugins=grpc:./usecase user.proto

	# compiling rest gateway
	# $(PROTOCMD) --proto_path=api \
	# -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	# --grpc-gateway_out=logtostderr=true,grpc_api_configuration=api/user.yaml:./usecase user.proto

	# compiling open api info
	# $(PROTOCMD) --proto_path=api \
	# -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	# --swagger_out=logtostderr=true:./api user.proto

	#remove "" on ubuntu machine
	#sed -i "" -e "s/RegisterUserServiceHandler/RegisterUserServiceMicroHandler/g" ./usecase/user.micro.go