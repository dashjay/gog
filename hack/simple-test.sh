#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
rootDir="$SCRIPT_DIR"/..

cd "$rootDir"

GOTOOLCHAIN=go1.18.10 go test ./... -v
GOTOOLCHAIN=go1.23.2 go test ./... -v