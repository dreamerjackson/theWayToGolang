#压力测试

go test -c # -c会生成可执行文件
# comment above and uncomment below to enable the race builder
# go test -c -race
PKG=$(basename $(pwd))  # 获取当前路径的最后一个名字，即为文件夹的名字
echo $PKG
while true ; do
        export GOMAXPROCS=$[ 1 + $[ RANDOM % 128 ]]
        ./$PKG.test $@ 2>&1   # $@代表可以加入参数   2>&1输出到控制台
done

##  pressure.sh
##  也可以加入参数：pressure.sh -test.v -test.run=ThisTestOnly