package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	content = `HTTP是如今互联网的基础协议，承载了互联网上的绝大部分应用层流量，并且从目前趋势来看，在未来10年，http仍然会是互联网应用的主要协议。Go语言自带“电池”，基于Go标准库我们可以轻松建立起一个http server处理客户端http请求，或创建一个http client向服务端发送http请求。

最初早期的http 1.0协议只支持短连接，即客户端每发送一个请求，就要和服务器端建立一个新TCP连接，请求处理完毕后，该连接将被拆除。显然每次tcp连接握手和拆除都将带来较大损耗，为了能充分利用已建立的连接，后来的http 1.0更新版和http 1.1支持在http请求头中加入Connection: keep-alive来告诉对方这个请求响应完成后不要关闭链接，下一次还要复用这个连接以继续传输后续请求和响应。后HTTP协议规范明确规定了HTTP/1.0版本如果想要保持长连接，需要在请求头中加上Connection: keep-alive，而HTTP/1.1版本将支持keep-alive长连接作为默认选项，有没有这个请求头都可以。

本文我们就来一起看看Go标准库中net/http包的http.Server和http.Client对keep-alive长连接的处理以及如何在Server和Client侧关闭keep-alive机制。

1. http包默认启用keep-alive
按照HTTP/1.1的规范，Go http包的http server和client的实现默认将所有连接视为长连接，无论这些连接上的初始请求是否带有Connection: keep-alive。

下面分别是使用go http包的默认机制实现的一个http client和一个http server：

默认开启keep-alive的http client实现：`
)

func pipeTest() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()

		scan := bufio.NewScanner(bytes.NewReader([]byte(content)))
		for scan.Scan() {
			line := scan.Text()
			_, _ = io.Copy(pw, bytes.NewReader([]byte(line)))
			time.Sleep(10 * time.Millisecond)
		}
	}()

	for {
		buff := make([]byte, 128)
		n, err := pr.Read(buff)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(string(buff[:n]))
		time.Sleep(20 * time.Millisecond)
	}
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-2]
	}
	return data
}

func main() {
	var mm sync.Map
	mm.Store("123", "321")
	fmt.Println(mm.Load("123"))
}
