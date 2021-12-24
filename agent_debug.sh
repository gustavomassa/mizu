#!/bin/bash

DLV=$(which dlv-dap)
if [ -x "${DLV}" ] ; then
	PATH="${GOPATH}/bin/:$PATH"
fi

sudo setcap cap_net_raw,cap_net_admin=eip ./agent/build/mizuagent
export GOGC=12800 
export NODE_NAME=dev

# The parameter -C 4 keeps the descriptor 3 opened
exec sudo -C 4 "$DLV" --only-same-user=false --headless=true --listen=:2345 --accept-multiclient exec ./agent/build/mizuagent -- \
        -i any \
        --tap \
        --api-server-address ws://localhost:8899/wsTapper & \
PID1=$! && \
read -r -d '' _ </dev/tty
sudo kill -9 $PID1