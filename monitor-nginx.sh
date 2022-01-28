#!/bin/bash

. /etc/profile

# check nginx process status, or could be Java process, like 'java -jar my.jar'
pgrep -f 'nginx' &>/dev/null 
if [ $? == 0 ]; then
    echo "nginx is running."
else
    /opt/wechat-alert -k my-robot-key -m "restarting nginx process" -mobile 172xxxx1234,173xxxx5678
	/usr/local/nginx
    echo "nginx restart."
fi