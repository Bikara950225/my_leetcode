package main

import (
	"bytes"
	"context"
	fmt "fmt"
	"io"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

func methodErr() (err error) {
	defer func() {
		fmt.Println(err.Error())
		err = fmt.Errorf("Hello")
	}()

	return fmt.Errorf("mock error")
}

func method1() {
	err := methodErr()
	e := fmt.Errorf("error: %w\n", err)
	fmt.Println(e.Error())
}

func start(addr [4]byte) {
	fd1, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd1)

	err = syscall.Bind(fd1, &syscall.SockaddrInet4{
		Port: 8080, Addr: addr,
	})
	if err != nil {
		panic(err)
	}
	err = syscall.Listen(fd1, 1024)
	if err != nil {
		panic(err)
	}

	for {
		cFd, _, err := syscall.Accept(fd1)
		if err != nil {
			panic(err)
		}

		func() {
			defer syscall.Close(cFd)

			buffer := make([]byte, 1024)
			n, err := syscall.Read(cFd, buffer)
			if err != nil {
				panic(err)
			}

			fmt.Printf("start(): [%d]%s\n", n, string(buffer[:n]))
		}()
	}
}

func start2(addr string) {
	listener, err := net.Listen("tcp", addr+":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		func() {
			defer conn.Close()

			bs, err := io.ReadAll(conn)
			if err != nil {
				panic(err)
			}

			fmt.Printf("start2(): %s\n", string(bs))
		}()
	}
}

type selfErr struct {
	count int
}

func (s selfErr) Error() string {
	s.count++
	return "123"
}

const bodyStr = `{
  "query": {},
  "fields": ["instanceId"],
  "page": 1,
  "page_size": 1
}`

func main() {
	selfDialContextFunc := func(ctx context.Context, network, addr string) (net.Conn, error) {
		dial := net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		return dial.Dial(network, addr)
	}

	httpCli := http.DefaultClient
	httpCli.Transport = &http.Transport{
		DialContext: selfDialContextFunc,
	}

	for range 3 {
		go func() {
			for {
				body := bytes.NewReader([]byte(bodyStr))
				url := fmt.Sprintf("http://sit.easyops.local:8079/v3/object/APP/instance/_search")
				req, err := http.NewRequest("POST", url, body)
				if err != nil {
					panic(err)
				}

				resp, err := httpCli.Do(req)
				if err != nil {
					panic(err)
				}
				io.Copy(os.Stdout, resp.Body)
				fmt.Println()
				resp.Body.Close()
				time.Sleep(2 * time.Second)
			}
		}()
	}
	var ch chan struct{}
	<-ch
}

type SessionStruct struct {
	Identity []string
	User     string
}

type Render struct {
	Session SessionStruct
	Prepare map[string]any
}

const templateContent = `{"filters": [{"operator": "{{.Session.User}}","target_id": {"$in": "{{.Prepare.appList.instanceId}}"}}]}`
