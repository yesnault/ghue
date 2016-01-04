#!/bin/bash

NAME="ghue"
for GOOS in windows darwin linux ; do
    for GOARCH in 386 amd64 arm; do
        if [[ $GOARCH == "arm" && $GOOS != "linux" ]]; then
          continue;
        fi;
        architecture="${GOOS}-${GOARCH}"
        echo "Building ${architecture} ${path}"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go build -ldflags "-X ${PROJECT_PATH}/${PROJECT_NAME}/cli/update.architecture=${architecture}" -o=bin/${NAME}-${architecture}
    done
done
