#!/bin/sh

. "$HOME/apps/hubx/src/aif.io/hubx/bin/Env.sh"

docker run -it -v $GOPATH:$GOPATH iseex.picp.io:30500/hubx/go-build:1.0 sh -c "$WORKROOT/build/make.sh"
