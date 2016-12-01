#!/usr/bin/env bash
set -e
set -o pipefail

port=8080
url="http://localhost:$port"

pushd cmd > /dev/null
	go build
	./cmd --port $port &
	id=$!
popd > /dev/null

pushd acceptance > /dev/null
	go test --url "$url" || true
popd > /dev/null

#Kill web server
kill -9 "$id"
