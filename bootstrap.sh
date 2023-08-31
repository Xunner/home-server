#! /usr/bin bash
currentDir=~/home-server
# 查找占用443端口的进程ID
pid=$(sudo lsof -t -i:443)

# 如果进程ID存在，发送SIGTERM信号给该进程
if [[ -n $pid ]]; then
    kill "$pid"
    echo "SIGTERM signal sent to process $pid."
else
    echo "No process is using port 443."
fi

nohup sudo go run $currentDir/main.go > output.log 2>&1 &
