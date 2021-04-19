# 初始化：
    go mod init xxx
# 配置包代理
    $ go env -w GO111MODULE=on
    $ go env -w GOPROXY=https://goproxy.cn,direct