#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
CONFIG_PATH=$(dirname $(dirname $CURDIR))/config
BinaryName=aster
echo "$CURDIR/bin/${BinaryName}"

exec $CURDIR/bin/${BinaryName} -config $CONFIG_PATH