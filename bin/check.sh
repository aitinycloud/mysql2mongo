#!/usr/bin/env bash
#

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

res=`ps -fe | grep $PROJECTPATH/bin/mysql2mongo | grep -v grep`
if [ $? -ne 0 ]
then
	echo "mysql2mongo stop"
	exit 1
else
	echo "mysql2mongo running"
	exit 0
fi
