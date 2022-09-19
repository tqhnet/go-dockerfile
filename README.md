# go-dockerfile
go容器打包的测试代码

### Docker命令大全

```
https://www.runoob.com/docker/docker-command-manual.html
```

### 功能

- 浏览器访问`0.0.0.0:8888`会输出文案
- 可将日志文件输出到本地路径

### 代码

代码

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8888",
	}
  fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello 路由器没有路!"))
}
```

dockerfile

```
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO_ENV=dev

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
RUN cp /build/app .

# 声明服务端口
EXPOSE 8888

# 启动容器时运行的命令
CMD ["/dist/app"]

```

### 打包容器

```
docker build . -t gotest
```

### 运行

```
docker run --name gotest -p 8888:8888 gotest
```

### 进入容器中

```
docker exec -it gotest /bin/sh
```

### 挂载目录运行

由于容器删除后里面的内容也会删除所以需要将日志文件等持久化挂载到外部

其中`/Users/macpro/Documents/xiaojing/go-test`是本地绝对路径，`:/app`是日志输出的路径

```
docker run --name gotest -p 8888:8888 -v /Users/macpro/Documents/xiaojing/go-test:/app -d gotest
```

### Docker命令

- **docker start** :启动一个或多个已经被停止的容器

- **docker stop** :停止一个运行中的容器

- **docker restart** :重启容器

```
docker start gotest
```

### Docker发布和部署

- 命令行登录账号：
  `docker login -u username`

- 新建一个tag，名字必须跟你注册账号一样
  `docker tag test:v1 username/test:v1`

- 推上去
  `docker push username/test:v1`

- 部署试下
  `docker run -dp 8080:8080 username/test:v1`

- 我的image

  `tqhnet/gotest`

