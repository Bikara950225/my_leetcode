package pipeline

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func method() {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	go func() {
		buffer := make([]byte, 5)
		i := 0
		for {
			n, err := r.Read(buffer)
			if err != nil {
				panic(err)
			}
			fmt.Printf("第%d次读取: %s\n", i, string(buffer[:n]))
			i++
		}
	}()

	j := 0
	for {
		time.Sleep(time.Second)

		js := strconv.Itoa(j)
		sb := strings.Builder{}
		for range 10 {
			sb.WriteString(js)
		}

		_, err := w.Write([]byte(sb.String()[:10]))
		if err != nil {
			panic(err)
		}
		j++
	}
}
