#!/usr/bin/env bash
#

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)


#echo "mysql2mongo start."
cd $PROJECTPATH
$PROJECTPATH/bin/mysql2mongo 2>&1 >> logs/info.log &
echo "*********************************"
echo "mysql2mongo process"
ps -fe | grep $PROJECTPATH/bin/mysql2mongo | grep -v grep
echo "*********************************"
sleep 1
