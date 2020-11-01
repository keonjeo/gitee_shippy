.PHONY: dep
dep:
	@echo "installing dependence..."
	@go get -v -x && go mod tidy

.PHONY: proto
proto:
	@echo "protocing..."
	@protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto

.PHONY: test
test: dep proto
	go test -v ./... -cover

.PHONY: build
build: proto
	@echo "building..."
	@go build -o user_srv .

.PHONY: run
run: proto
	@echo "running..."
	@go run main.go

.PHONY: docker_build
docker_build: dep proto
	@docker build -t user_srv:latest .

.PHONY: docker_run
docker_run:
	@docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user_srv
