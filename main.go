package main

import (
	"fmt"
	"os"

	"github.com/sm-2020/tracer/src/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
