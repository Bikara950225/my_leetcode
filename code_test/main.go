package main

import (
	"encoding/json"
	fmt "fmt"
	"io"
	"net"
	"syscall"
	"unsafe"
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
	//rf, err := os.Open("/tmp/test1.log")
	//if err != nil {
	//	panic(err)
	//}
	//defer rf.Close()
	//
	//iFd := rf.Fd()
	//
	//wf, err := os.OpenFile("/tmp/test2.log", os.O_CREATE|os.O_WRONLY, 0777)
	//if err != nil {
	//	panic(err)
	//}
	//defer wf.Close()
	//wFd := wf.Fd()
	//
	//var ofs int64 = 0
	//_, err = syscall.Sendfile(int(wFd), int(iFd), &ofs, 3)
	//if err != nil {
	//	panic(err)
	//}
	var i1 int8 = 100
	var i2 int8 = 127
	i1p := unsafe.Pointer(&i1)
	i2p := unsafe.Pointer(uintptr(i1p) + 1)
	//cpi2 := *(*int8)(unsafe.Pointer(i2p))
	fmt.Println(unsafe.Pointer(&i2))
	fmt.Println(i2p)
	fmt.Println(*(*int8)(i2p))
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
