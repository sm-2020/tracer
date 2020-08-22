package main

import (
	"fmt"
	"os"
	"github.com/sm-2020/tracer/src/cmd"
)

func main() {
	if err := tracercli.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
