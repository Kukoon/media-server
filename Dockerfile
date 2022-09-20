##
# Compile application
##
FROM golang:alpine AS build-env
ARG VERSION
WORKDIR /app
COPY . .
# Build Docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseDepth 4 --parseInternal -g web/docs.go -o web/docs
# ge dependencies
RUN go mod tidy
# build binary
RUN CGO_ENABLED=0 go build -ldflags="-w -s -X main.VERSION=$VERSION" -o server


##
# Build Image
##
FROM scratch
COPY --from=build-env ["/etc/ssl/cert.pem", "/etc/ssl/certs/ca-certificates.crt"]
COPY --from=build-env /app/server /server

WORKDIR /
ENTRYPOINT ["/server", "-c", ""]
