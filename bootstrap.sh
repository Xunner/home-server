#! /usr/bin bash
cd ~/home-server || exit 3
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

# build
go build

if [ ! -d "output" ]; then
  mkdir output
fi
NOW=$(date '+%Y-%m-%d_%H_%M_%S')

# run
nohup sudo bash home-server > output/output_"$NOW".log 2>&1 &
cat output/output_"$NOW".log
