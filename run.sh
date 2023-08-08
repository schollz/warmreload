#!/bin/bash
ps -ef | grep 'warmreload' | grep -v grep | grep -v run | awk '{print $2}' | xargs -r kill -9
cd "$(dirname "$0")"
nohup ./warmreload -path /home/we/dust/code/ >/dev/null 2>&1 &
