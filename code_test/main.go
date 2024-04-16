package main

import (
	"encoding/json"
	fmt "fmt"
	"io"
	"net"
	"syscall"
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

type ss struct {
	num int
}

func (s *ss) method17() {
	fmt.Println(s.num)
}

type Node struct {
	Val any
	Sub []*Node
}

type Tree struct {
	Root *Node
}

type MM map[string]any

type ParseText struct {
	MM
	Name string
	Age  int
}

type selfSlice[T any] []T

func (s selfSlice[T]) String() string {
	bs, _ := json.Marshal(s)
	sws := string(bs)
	return sws
}

type selfSliceString []string

func (s selfSliceString) String() string {
	bs, _ := json.Marshal(s)
	return string(bs)
}

func main() {
	for i := range 4 {
		fmt.Println(i)
	}
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
