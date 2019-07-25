#!/usr/bin/env bash
#

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

echo "mysql2mongo restart."
$PROJECTPATH/bin/stop.sh
sleep 1
$PROJECTPATH/bin/start.sh
