#!/bin/sh

. "$HOME/apps/hubx/src/aif.io/hubx/bin/Env.sh"

go build -ldflags "-w -s" -o $WORKROOT/build/$EXE_NAME $WORKROOT/cmd/main.go
echo "success build exe to ->  $WORKROOT/build/$EXE_NAME"
exit 0
