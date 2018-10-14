FROM golang:1-stretch

RUN apt-get update -y && apt-get install -y upx-ucl

WORKDIR /opt/status_service
COPY status_service.go .
RUN GOARCH=386 GOOS=linux go build -ldflags "-s" \
&& upx --best status_service

FROM scratch
COPY --from=0 /opt/status_service /
ENTRYPOINT [ "/status_service" ]
