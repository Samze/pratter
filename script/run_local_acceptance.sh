#!/usr/bin/env bash
set -e
set -o pipefail

port=8080
url="http://localhost:$port"

go build
./pratter --port $port &
id=$!

pushd acceptance
	go test --url "$url" || true
popd

#Kill web server
kill -9 "$id"
