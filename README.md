# go-wechat-robot-sender

You can use this to send msg to wechat(enterprise) msg

can be used as System monitor with shell-script

**Only support text msg.**

like monitor-nginx.sh

```shell
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
```

then add it to crontab

```bash
crontab -u root -e

...
*/1 * * * * /opt/monitor-nginx.sh
```





## RequireMents

-   Go >= 1.16



## Building project

Downloading `git clone https://github.com/RRRRIC/go-wechat-rebot-sender.git`

### Under Linux/Mac Os

```bash
# Buildingh project

# Linux 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build

# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```



### Under Windows Os

```bash
# Mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
 
# Linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build

# Windows
go build
```



## Usage

```bash
# Only do this the first time
chmod +7 wechat-alert

./wechat-alert -help

./wechat-alert -k my-rebot-key -m my-msg -mobile 139xxxx1234,158xxxx2345,172xxxx3456
```

