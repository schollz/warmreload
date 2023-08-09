#!/bin/bash
ps -ef | grep 'warmreload' | grep -v grep | grep -v run | awk '{print $2}' | xargs -r kill -9
cd /home/we/dust/code/warmreload
make warmreload
nohup ./warmreload -path /home/we/dust/code/ >/dev/null 2>&1 &
