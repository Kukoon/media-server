##
# Compile application
##
FROM golang:latest AS build-env
WORKDIR /app
COPY . .
# Build Docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseDepth 4 --parseInternal -g web/docs.go -o web/docs
# ge dependencies
RUN go mod tidy
# build binary
RUN CGO_ENABLED=0 go build -ldflags="-w -s -X main.VERSION=$CI_COMMIT_REF_SLUG" -o server


##
# Build Image
##
FROM scratch
COPY --from=build-env /app/server /server
COPY --from=build-env /app/config_example.toml /config.toml

WORKDIR /
ENTRYPOINT ["/server"]
