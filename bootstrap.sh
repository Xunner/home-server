#! /usr/bin bash

# 查找占用443端口的进程ID
PREV_PID=$(sudo lsof -t -i:443)

# 如果进程ID存在，发送SIGTERM信号给该进程
if [[ -n $PREV_PID ]]; then
  OUTPUT=$(sudo kill "$PREV_PID")
  if [ $? != 0 ]; then
    echo "Prev program is running, killing failed: $OUTPUT"
    exit 1
  fi
  echo "Prev program killed"
fi

if [ ! -d "output" ]; then
  mkdir output
fi
CURRENT_DIR=~/home-server
NOW=$(date '+%Y-%m-%d_%H:%M:%S')

nohup sudo go run $CURRENT_DIR/main.go > output/output_"$NOW".log 2>&1 &
cat output/output_"$NOW".log
