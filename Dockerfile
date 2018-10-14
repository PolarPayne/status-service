# --- build ---
FROM golang:1-stretch AS build

RUN apt-get update -y && apt-get install -y upx-ucl

WORKDIR /opt/status_service
COPY . .

RUN make all

# --- final ---
FROM scratch AS final
COPY --from=build /opt/status_service/status-service-min /status-service
ENTRYPOINT [ "/status-service" ]
