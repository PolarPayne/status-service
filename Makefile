.PHONY: all
all: test status-service status-service-min

status-service:
	@echo "+ go build"
	@GOARCH=386 GOOS=linux go build -ldflags "-s" -o status-service

status-service-min: status-service
	@echo "+ minify"
	@upx --best status-service -o status-service-min

.PHONY: docker-build
docker-build:
	docker build -t status-service .

.PHONY: test
test:
	@echo "+ go test"
	@go test .

.PHONY: clean
clean:
	rm -f status-service status-service-min
