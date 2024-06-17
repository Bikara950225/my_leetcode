package main

import (
	_ "embed"
	fmt "fmt"
	"github.com/google/btree"
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

type selfErr struct {
	count int
}

func (s selfErr) Error() string {
	s.count++
	return "123"
}

var helloFile []byte

type entry struct {
	id  int
	val string
}

func main() {
	bTree := btree.NewG(10, func(a, b *entry) bool {
		return a.id < b.id
	})
	for i := range 100 {
		bTree.ReplaceOrInsert(&entry{
			id: i, val: fmt.Sprintf("val%d", i),
		})
	}

	fmt.Println(bTree.Len())
	fmt.Println(bTree.Get(&entry{id: 50}))
	bTree.AscendRange(&entry{id: 49}, &entry{id: 100}, func(item *entry) bool {
		fmt.Println(item)
		return true
	})
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
