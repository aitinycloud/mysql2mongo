# mysql2mongo.

mysql2mongo for sync mysql master database to mongodb, implement by golang.

Befor build please install task tool. go version >= 1.10.4 && git version >= 2.1.1 go get -v github.com/go-task/task/cmd/task

task -l

eg:
    task prebuild

    task build

    task run

    task stop


NOTE: 

In the development state.


After build.

mysql2mongo Start

./bin/start.sh

mysql2mongo Stop

./bin/stop.sh

mysql2mongo Restart

./bin/restart.sh

mysql2mongo status

./bin/check.sh

