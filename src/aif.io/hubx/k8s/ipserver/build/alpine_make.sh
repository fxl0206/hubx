#!/bin/sh

. "$HOME/apps/hubx/src/aif.io/hubx/bin/Env.sh"

docker run --net host -v $GOPATH:$GOPATH iseex.picp.io:30500/hubx/go-build:1.0 $WORKROOT/build/make.sh
