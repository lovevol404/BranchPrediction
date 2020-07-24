# BranchPrediction
分支预测测试代码,下载对应平台编译好的二进制可执行文件运行即可，观察输出结果即可得出是否支持分支预测。
如果第二行的耗时与第一行耗时差距比较大，就说明可能是分支预测导致的性能提升
# 构建命令
go build -o bin/BranchPrediction_mac 

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/BranchPrediction_linux 

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/BranchPrediction_windows.exe 

