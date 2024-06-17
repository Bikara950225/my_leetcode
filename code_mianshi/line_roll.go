package code_mianshi

import (
	"fmt"
	"time"
)

func PrintMsg(msg string) {
	for {
		for _, r := range "|\\-/" {
			fmt.Printf("\r%c: %s", r, msg)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
