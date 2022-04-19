package main

import (
	"bufio"
	"fmt"
	"gateserver/applications"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	app := applications.NewApplicationInstance()
	app.Start(index, os.Args[2])

	exitchan := make(chan struct{})
	go func(ch chan<- struct{}) {
		reader := bufio.NewReader(os.Stdin)
		for {
			str, _, err := reader.ReadLine()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal: console read line, detail: %s", str)
				break
			}

			cmd := string(str)
			if strings.EqualFold(cmd, "quit") || cmd == "2" {
				close(ch)
				break
			}
		}
	}(exitchan)

	for !app.IsClosed() {
		select {
		case <-exitchan:
			app.Stop()
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
