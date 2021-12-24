#!/bin/bash

sudo setcap cap_net_raw,cap_net_admin=eip ./agent/build/mizuagent
sudo GOGC=12800 NODE_NAME=dev \
    ./agent/build/mizuagent \
        -i any \
        --tap \
        --api-server-address ws://localhost:8899/wsTapper