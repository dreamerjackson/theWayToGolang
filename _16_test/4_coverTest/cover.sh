
# 创建临时文件，并测试代码覆盖率
cover() {
    local t=$(mktemp -t cover)
    go test $COVERFLAGS -coverprofile=$t $@ \
        && go tool cover -func=$t \
        && unlink $t
}
# 创建临时文件，并测试代码覆盖率，在web中显示
coverweb() {
    local t=$(mktemp -t cover)
    go test $COVERFLAGS -coverprofile=$t $@ \
        && go tool cover -html=$t \
        && unlink $t
}

## cover bytes 可以加一个参数，是package名字
cover
coverweb


## 启动
# ./cover.sh
