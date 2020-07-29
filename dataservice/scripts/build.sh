#/bin/sh

set -e 

APP_NAME=${OUT_FILE_NAME:-app}
APP_DIR="${1:-./dataservice}"

env GOMOD=${APP_DIR}

export CGO_ENABLED=0 
# export GOOS=linux 
# export GOARCH=amd64
go test ${APP_DIR}
env GOOS=linux GOARCH=amd64 GOFLAGS=-insecure go build -v -installsuffix cgo -o ${APP_NAME} ${APP_DIR}
#env GOOS=linux GOARCH=amd64 go install -v  -installsuffix cgo -o ${APP_NAME} ${APP_DIR}