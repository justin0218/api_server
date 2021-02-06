echo 更新代码

git fetch --all

git reset --hard origin/master

echo 开始构建

export GOPROXY=https://goproxy.cn

/usr/local/go/bin/go build -o api_server cmd/main.go

echo 重启服务

supervisorctl restart api_server