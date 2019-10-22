#!/bin/bash

APP_DIR="/go/src/github.com/${GITHUB_REPOSITORY}/"

# shellcheck disable=SC2164
mkdir -p "${APP_DIR}" && cp -r ./ "${APP_DIR}" && cd "${APP_DIR}"

export GO111MODULE=on
go mod tidy
go mod verify

if [[ "$1" == "lint" ]]; then
    echo "############################"
    echo "# Running GolangCI-Lint... #"
    echo "############################"
    golangci-lint --version
    echo
    golangci-lint run --tests --disable-all --enable=goimports --enable=golint --enable=govet --enable=errcheck ./...
fi