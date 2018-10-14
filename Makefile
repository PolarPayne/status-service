.PHONY: all
all: test status-service status-service-min

status-service: status_service.go
	@echo "+ go build"
	@GOARCH=386 GOOS=linux go build -ldflags "-s" -o status-service

status-service-min: status-service
	@echo "+ minify"
	@rm -f status-service-min
	@upx --best status-service -o status-service-min

.PHONY: docker-build
docker-build:
	docker build -t status-service .

.PHONY: test
test: status_service.go status_service_test.go
	@echo "+ go test"
	@go test .

.PHONY: clean
clean:
	rm -f status-service status-service-min
