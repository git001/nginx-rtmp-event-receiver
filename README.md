## nginx-rtmp-event-receiver

[![Sourcegraph](https://sourcegraph.com/github.com/git001/nginx-rtmp-event-receiver/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/git001/nginx-rtmp-event-receiver?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/git001/nginx-rtmp-event-receiver?style=flat-square)](https://goreportcard.com/report/github.com/git001/nginx-rtmp-event-receiver)
[![Twitter](https://img.shields.io/badge/twitter-@me2digital.svg?style=flat-square)](https://twitter.com/me2digital)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/git001/nginx-rtmp-event-receiver/main/LICENSE)


This small too should help to create some poor man access control for the 
https://github.com/sergey-dryabzhinsky/nginx-rtmp-module which is based on 
https://github.com/arut/nginx-rtmp-module

I use the [gofiber](https://github.com/gofiber/fiber/) framework to build this tool.

### Build

```
mkdir bin
go build -o bin/event-receiver_with_accesslist
```

### configure

#### event-receiver

```yaml
---
# myconfig.yaml
config:
  # show debug messages
  debug: false
  # 256 * 1024
  Concurrency: 262144
  # disable fiber startup messsage
  DisableStartupMessage: false
  ReadTimeout: 2s
  WriteTimeout: 2s
  IdleTimeout: 5s
  ListenPort: ":3000"

streams:
  # the app name in the rtmp payload
  stream-3412: true
```

```shell
#.env
MY_CONFIG=myconfig.yaml
```

##### tshark output

You can check what you get in the app field with this tshark command

```shell
tshark -f 'port 3000' -d tcp.port==3000,http -i lo  -Y "http.request.method==POST" -Tfields -e text

Timestamps,POST /on_connect HTTP/1.0\r\n,\r\n,
Form item: "camid" = "xx",
Form item: "token" = "xxx",
Form item: "call" = "connect",
Form item: "app" = "camera01",
Form item: "flashver" = "FMLE/3.0 (compatible; Lavf58.59",
Form item: "swfurl" = "",
Form item: "tcurl" = "rtmps://subdomain.domain.com:1936/camera01?camid=xx&token=xxx",
Form item: "pageurl" = "",
Form item: "epoch" = "145583460",
Form item: "&addr" = "xxx.xxx.xxx.xxx",
Form item: "clientid" = "xxx"
```

#### nginx-rtmp-module

```
...
on_connect        http://127.0.0.1:3000/on_connect;
on_play           http://127.0.0.1:3000/on_play;
on_playlist       http://127.0.0.1:3000/on_playlist;
on_publish        http://127.0.0.1:3000/on_publish;
on_done           http://127.0.0.1:3000/on_done;
on_play_done      http://127.0.0.1:3000/on_play_done;
on_publish_done   http://127.0.0.1:3000/on_publish_done;
on_record_started http://127.0.0.1:3000/on_record_started;
on_record_done    http://127.0.0.1:3000/on_record_done;
on_update         http://127.0.0.1:3000/on_update;
...
```
