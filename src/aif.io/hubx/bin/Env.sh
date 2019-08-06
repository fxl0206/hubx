#!/bin/sh

PRG="$0"
while [ -h "$PRG" ] ; do
  ls=`ls -ld "$PRG"`
  link=`expr "$ls" : '.*-> \(.*\)$'`
  if expr "$link" : '/.*' > /dev/null; then
     PRG="$link"
  else
     PRG=`dirname "$PRG"`/"$link"
  fi
done
cd `dirname "$PRG"`
export WORKROOT=`pwd | xargs dirname`
export GOPATH=$HOME/apps/hubx
export EXE_NAME=hubx
echo "use WORKROOT : $WORKROOT"
echo "use GOPATH   : $GOPATH"
echo "use EXE_NAME : $EXE_NAME"
