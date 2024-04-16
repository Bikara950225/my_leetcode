package main

import (
	"encoding/json"
	fmt "fmt"
	"io"
	"net"
	"reflect"
	"strconv"
	"strings"
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

func strToIp(ip string) uint32 {
	ipSplit := strings.Split(ip, ".")
	if len(ipSplit) != 4 {
		panic("ff")
	}
	var ipInt uint32
	for i, subIpStr := range ipSplit {
		subIp, _ := strconv.ParseUint(subIpStr, 10, 32)
		ipInt |= uint32(subIp) << (24 - 8*i)
	}
	return ipInt
}

func ipToStr(ip uint32) string {
	byteH := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&ip)),
		Len:  4, Cap: 4,
	}
	bs := (*[]byte)(unsafe.Pointer(&byteH))
	sb := strings.Builder{}
	for i := len(*bs) - 1; i >= 0; i-- {
		subIp := strconv.Itoa(int((*bs)[i]))
		sb.WriteString(subIp)
		if i != 0 {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func main() {
	ip1 := "172.30.0.57"
	ip2 := "173.30.0.58"

	fmt.Println(strToIp(ip1))
	fmt.Println(strToIp(ip2))

	fmt.Println(ipToStr(strToIp(ip1)))
	fmt.Println(ipToStr(strToIp(ip2)))
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
