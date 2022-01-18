package main

import (
	"fmt"
	"time"

	"github.com/chentanyi/awake"
)

var (
	progresses = []byte{'|', '/', '-', '\\'}
)

func main() {
	go awake.Awake()
	for {
		for _, progress := range progresses {
			fmt.Printf("\r%c Awaking the desktop.", progress)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
