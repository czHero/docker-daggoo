// main.go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"os/exec"
	"syscall"
)

func main() {
	var cmd *exec.Cmd
	var res []byte
	var err error
	fmt.Println("Program started...")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	s := <-ch

	if s == syscall.SIGTERM {
		fmt.Println("SIGTERM received!")
		cmd = exec.Command("/etc/init.d/celeryd", "stop")
		if res, err = cmd.Output(); err != nil {
	        fmt.Println(err)
	        os.Exit(1)
	    }
	    fmt.Println(string(res))
	}

	fmt.Println("Exiting...")
}