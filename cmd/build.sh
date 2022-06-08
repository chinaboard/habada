#!/usr/bin/env bash

#set -x

GitStatus=`git status -s`
BuildTime=`date +'%Y.%m.%d.%H%M%S'`
BuildGoVersion=`go version`

LDFlags=" \
    -X 'github.com/chinaboard/habada/pkg/bininfo.GitStatus=${GitStatus}' \
    -X 'github.com/chinaboard/habada/pkg/bininfo.BuildTime=${BuildTime}' \
    -X 'github.com/chinaboard/habada/pkg/bininfo.BuildGoVersion=${BuildGoVersion}' \
"

ROOT_DIR=`pwd`

# 如果可执行程序输出目录不存在，则创建
if [ ! -d ${ROOT_DIR}/bin ]; then
  mkdir bin
fi

# 编译多个可执行程序
cd ${ROOT_DIR} && go build -ldflags "$LDFlags" -o ${ROOT_DIR}/bin/habada-${GOOS}-${GOARCH}

ls -lrt ${ROOT_DIR}/bin &&
echo 'build done.'