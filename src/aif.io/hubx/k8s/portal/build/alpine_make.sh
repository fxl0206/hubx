#!/bin/sh

. "$HOME/apps/hubx/src/aif.io/hubx/bin/Env.sh"

docker run -it -v $GOPATH:$GOPATH doc.hubx.site/hubx/go-build:1.0 sh -c "$WORKROOT/build/make.sh"
