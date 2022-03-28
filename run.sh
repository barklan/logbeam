#!/usr/bin/env bash

set -eo pipefail

DC="${DC:-exec}"

# If we're running in CI we need to disable TTY allocation for docker-compose
# commands that enable it by default, such as exec and run.
TTY=""
if [[ ! -t 1 ]]; then
    TTY="-T"
fi

# -----------------------------------------------------------------------------
# Helper functions start with _ and aren't listed in this script's help menu.
# -----------------------------------------------------------------------------

function _dc {
    export DOCKER_BUILDKIT=1
    docker-compose ${TTY} "${@}"
}

function _use_env {
    set -o allexport; . .env; set +o allexport
}

# ----------------------------------------------------------------------------

up() {
    reflex -c reflex.conf --decoration=fancy
}

up:logdip() {
    _use_env
    go run ./cmd/logdip "${@}"
}

build:logdip() {
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOGC=off go build \
    -ldflags='-w -s -extldflags "-static"' -a -o ./.cache/logdip/logdip ./cmd/logdip/.
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOGC=off go build \
    -ldflags='-w -s -extldflags "-static"' -a -o ./.cache/logdip/logdip.exe ./cmd/logdip/.
}

up:compose() {
    docker-compose --profile main build --parallel
    docker-compose --profile main up
}

lint() {
    golangci-lint run --enable-all  --disable=wsl,varnamelen,testpackage,gomnd,exhaustivestruct
}

test() {
    go test -v -race -cover ./...
}

test:e2e() {
    go test -v -race -tags e2e ./...
}

env:prod() {
    ENV_FOR=prod bash ./scripts/build_env.sh
}

up:docs() {
    docker run -it --rm -p 8080:80 \
    -v "$(pwd)"/docs:/usr/share/nginx/html/swagger/ \
    -e SPEC_URL=swagger/openapi.yml redocly/redoc
}

docs:gen() {
    docker run --user 1000:1000 --rm -v "$(pwd)"/docs:/spec redocly/openapi-cli bundle -o bundle.json --ext json openapi.yml
    docker run --rm -it --ulimit nofile=122880:122880 -m 3G \
    -v "${PWD}"/docs:/docs -w /docs swaggerapi/swagger-codegen-cli-v3 generate -i https://raw.githubusercontent.com/barklan/logdip/main/docs/openapi.yml -l go -o ./go
}

# -----------------------------------------------------------------------------

function help {
    printf "%s <task> [args]\n\nTasks:\n" "${0}"

    compgen -A function | grep -v "^_" | cat -n
}

TIMEFORMAT=$'\nTask completed in %3lR'
time "${@:-help}"
