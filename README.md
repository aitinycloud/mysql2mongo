# mysql2mongo tools.

mysql2mongo for sync master database to slaves database, implement by golang.

Befor build please install task tool.
go version >= 1.10.4 && git version >= 2.1.1
go get -v github.com/go-task/task/cmd/task

task -l
eg:
    task prebuild
    task build
    task run
    task stop

NOTE: 
golang.org/x/XXX 提示错误,从 https://github.com/golang  下载,解压到GOPATH目录中.
go.etcd.io/etcd 提示错误,从 https://github.com/etcd-io/etcd  下载,解压到GOPATH目录中.

After build.
# mysql2mongo Start
cd mysql2mongo PATH.
./bin/start.sh

# mysql2mongo Stop
cd mysql2mongo PATH.
./bin/stop.sh

# mysql2mongo Restart
cd mysql2mongo PATH.
./bin/restart.sh

# mysql2mongo status
cd mysql2mongo PATH.
./bin/check.sh


# email: 
