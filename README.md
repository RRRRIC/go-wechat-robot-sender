# go-wechat-rebot-sender

You can use this to send msg to wechat(enterprise) msg
can be used as System monitor with shell-script

**Only support text msg.**



## RequireMent

-   Go >= 1.16



## Building



### Under Linux/Mac Os

```bash
# Buildingh project

# Linux 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build filename.go

# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build filename.go
```



### User Windows Os

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
```



## Usage

```bash
# Only do this the first time
chmod +7 wechat-alert

./wechat-alert -k my-rebot-key -m my-msg

# -m Msg has default value : test-msg
# -k Key has not default value

```


