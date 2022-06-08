#!/usr/bin/env bash

set -x

export GOOS=darwin
export GOARCH=arm64
sh cmd/build.sh